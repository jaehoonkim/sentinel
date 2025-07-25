package fetcher

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/panta/machineid"

	"github.com/jaehoonkim/sentinel/pkg/agent/log"
	"github.com/jaehoonkim/sentinel/pkg/agent/sentinel"
	"github.com/jaehoonkim/sentinel/pkg/agent/scheduler"
	"github.com/jaehoonkim/sentinel/pkg/agent/service"
	sessionv1 "github.com/jaehoonkim/sentinel/pkg/manager/model/session/v1"
)

const (
	defaultPollingInterval = 5 // * time.Second
	minPollingInterval     = 5
)

type Fetcher struct {
	bearerToken     string
	machineID       string
	clusterId       string
	sentinelAPI     *sentinel.SynapseAPI
	ticker          *time.Ticker
	pollingInterval int
	scheduler       *scheduler.Scheduler
	done            chan struct{}
}

func NewFetcher(bearerToken, manager, clusterId string, scheduler *scheduler.Scheduler) (*Fetcher, error) {
	id, err := machineid.ID()
	if err != nil {
		return nil, err
	}
	id = strings.ReplaceAll(id, "-", "")

	api, err := sentinel.NewSynapseAPI(manager)
	log.Debugf("api in fetcher.go : %s\n", *api)
	if err != nil {
		return nil, err
	}

	return &Fetcher{
		bearerToken:     bearerToken,
		machineID:       id,
		clusterId:       clusterId,
		sentinelAPI:     api,
		ticker:          time.NewTicker(defaultPollingInterval * time.Second),
		pollingInterval: defaultPollingInterval,
		scheduler:       scheduler,
		done:            make(chan struct{})}, nil
}

func (f *Fetcher) ChangePollingInterval(interval int) (int, error) {
	if interval < minPollingInterval {
		return 0, fmt.Errorf("interval(%d) you want to change is less than the minimum interval(%d)", interval, minPollingInterval)
	}

	if f.pollingInterval == interval {
		return 0, nil
	}
	f.pollingInterval = interval
	f.ticker.Reset(time.Second * time.Duration(interval))

	return interval, nil
}

func (f *Fetcher) Done() <-chan struct{} {
	return f.done
}

func (f *Fetcher) Cancel() {
	close(f.done)
}

func (f *Fetcher) Polling(ctx context.Context) error {
	if f == nil || f.ticker == nil {
		return fmt.Errorf("fetcher or fetcher.ticker is not created")
	}

	go f.UpdateServiceProcess()

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Debugf("polling context done")
				return
			case <-f.ticker.C:
				f.poll()
			}
		}
	}()

	return nil
}

func (f *Fetcher) poll() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// get services from manager
	respData, err := f.sentinelAPI.GetServices(ctx)
	if err != nil {
		log.Errorf("Failed to polling: error: %s\n", err.Error())

		// if session token is expired, retry handshake
		if f.sentinelAPI.IsTokenExpired() {
			f.ticker.Stop()
			f.RetryHandshake()
		}

		return
	}

	f.ChangeAgentConfigFromToken()

	log.Debugf("Received %d service from manager.", len(respData))

	if len(respData) == 0 {
		<-time.After(time.Duration(f.pollingInterval) * time.Second)
		return
	}

	// respData -> services
	recvServices, failed := service.ConvertServiceListManagerToAgent(respData)

	if len(failed) > 0 {
		log.Debugf("Failed to convert %d service\n", len(failed))
		f.UpdateFailedToConvertServices(failed)
	}

	// catch sentinelagent service
	if ok := f.CatchSynapseAgentService(recvServices); ok {
		return
	}

	// Register new services.
	f.scheduler.RegisterServices(recvServices)
}

func (f *Fetcher) ChangeAgentConfigFromToken() {
	claims := new(sessionv1.ClientSessionPayload)
	jwt_token, _, err := jwt.NewParser().ParseUnverified(f.sentinelAPI.GetToken(), claims)
	if _, ok := jwt_token.Claims.(*sessionv1.ClientSessionPayload); !ok || err != nil {
		log.Warnf("Failed to bind payload : %v\n", err)
		return
	}

	if interval, err := f.ChangePollingInterval(claims.PollInterval); err != nil {
		log.Warnf("Failed to change polling interval : %v\n", err)
	} else {
		if interval != 0 {
			log.Debugf("Change polling interval to %v\n", claims.PollInterval)
		}
	}

	if log.GetLogger().GetLevel() != strings.ToLower(claims.Loglevel) {
		if err := log.GetLogger().SetLevel(claims.Loglevel); err == nil {
			log.Debugf("Changed logger level to %s\n", claims.Loglevel)
		}
	}
}

func (f *Fetcher) UpdateServiceProcess() {
	for update := range f.scheduler.NotifyServiceUpdate() {
		<-time.After(time.Millisecond * 100)

		go func(up service.ServiceUpdateInterface) {
			serv := service.ConvertServiceStepUpdateAgentToManager(up)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := f.sentinelAPI.UpdateServices(ctx, serv); err != nil {
				switch serv.Version {
				case "v3":
					log.Errorf("Failed to update service on manager : service_uuid:%s, error:%s\n", serv.V3.Uuid, err.Error())
				case "v4":
					log.Errorf("Failed to update service on manager : service_uuid:%s, error:%s\n", serv.V4.Uuid, err.Error())
				}
			}

			f.scheduler.UpdateServiceStatus(up)
		}(update)
	}
}

func (f *Fetcher) CatchSynapseAgentService(services map[string]service.ServiceInterface) bool {
	exist := false

	for _, svc := range services {
		switch ver := svc.Version(); ver {
		case service.SERVICE_VERSION_V1:
			svcv1 := svc.(*service.ServiceV1)
			for _, step := range svcv1.Steps {
				if step.Command != nil {
					method := step.Command.Method

					switch method {
					case "sentinel.agent_pod.rebounce":
						exist = true
						f.RebounceAgentPod(ver, svcv1.Id)
					case "sentinel.agent.upgrade":
						exist = true
						f.UpgradeAgent(ver, svcv1.Id, step.Command.Args)
					}
					if exist {
						return exist
					}
				}
			}
		case service.SERVICE_VERSION_V2:
			svcv2 := svc.(*service.ServiceV2)
			for _, step := range svcv2.Flow {
				if step.Command != "" {
					method := step.Command

					switch method {
					case "sentinel.agent_pod.rebounce":
						exist = true
						f.RebounceAgentPod(ver, svcv2.Id)
					case "sentinel.agent.upgrade":
						exist = true
						f.UpgradeAgent(ver, svcv2.Id, step.Inputs.GetInputs())
					}
					if exist {
						return exist
					}
				}
			}
		}
	}

	return exist
}

func (f *Fetcher) RemainServices() map[string]service.ServiceStatus {
	return f.scheduler.CleanupRemainingServices()
}

func (f *Fetcher) UpdateFailedToConvertServices(failed []service.FailedConvertService) {
	for _, d := range failed {
		var up service.ServiceUpdateInterface
		switch d.Data.Version {
		case "v3":
			dd := d.Data.V3
			t := time.Now()
			up = &service.UpdateServiceV1{
				Uuid:      dd.Uuid,
				StepCount: len(dd.Steps),
				Sequence:  0,
				Status:    service.StepStatusFail,
				Result:    d.Err.Error(),
				Started:   t,
				Ended:     t,
			}
		case "v4":
			dd := d.Data.V4
			t := time.Now()
			up = &service.UpdateServiceV2{
				Id:        dd.Uuid,
				StepCount: dd.StepMax,
				Sequence:  0,
				Status:    service.StepStatusFail,
				Result:    d.Err.Error(),
				Started:   t,
				Ended:     t,
			}
		}

		serv := service.ConvertServiceStepUpdateAgentToManager(up)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := f.sentinelAPI.UpdateServices(ctx, serv); err != nil {
			switch serv.Version {
			case "v3":
				log.Errorf("Failed to update service on manager : service_uuid:%s, error:%s\n", serv.V3.Uuid, err.Error())
			case "v4":
				log.Errorf("Failed to update service on manager : service_uuid:%s, error:%s\n", serv.V4.Uuid, err.Error())
			}
		}
	}
}
