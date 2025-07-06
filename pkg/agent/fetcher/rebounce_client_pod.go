package fetcher

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/jaehoonkim/sentinel/pkg/agent/log"
	"github.com/jaehoonkim/sentinel/pkg/agent/service"
)

func (f *Fetcher) RebounceAgentPod(version service.Version, serviceId string) {
	t := time.Now()
	log.Debugf("SynapseAgentPod Rebounce: start")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	up := service.CreateUpdateService(version, serviceId, 1, 0, service.StepStatusProcessing, "", t, time.Time{})
	if err := f.sentinelAPI.UpdateServices(ctx, service.ConvertServiceStepUpdateAgentToManager(up)); err != nil {
		log.Errorf("SynapseAgentPod Rebounce: failed to update service status(processing): error: %s\n", err.Error())
	}

	f.ticker.Stop()
	log.Debugf("SynapseAgentPod Rebounce: polling stop")

	done := make(chan struct{})
	go func() {
		defer func() {
			done <- struct{}{}
		}()

		log.Debugf("SynapseAgentPod Rebounce: waiting to process the remaining services(timeout:30s)")

		for {
			<-time.After(time.Second * 3)
			services := f.RemainServices()
			if len(services) == 0 {
				break
			}

			buf := bytes.Buffer{}
			buf.WriteString("SynapseAgentPod Rebounce: waiting remain services:")
			for uuid, status := range services {
				buf.WriteString(fmt.Sprintf("\n\tuuid: %s, status: %s", uuid, status.String()))
			}
			log.Debugf(buf.String() + "\n")
		}
	}()

	select {
	case <-time.After(time.Second * 30):
		log.Debugf("SynapseAgentPod Rebounce: waiting timeout")
	case <-done:
		log.Debugf("SynapseAgentPod Rebounce: waiting done")
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel2()

	up = service.CreateUpdateService(version, serviceId, 1, 0, service.StepStatusSuccess, "SynapseAgent pod rebounce will be complete", t, time.Now())
	if err := f.sentinelAPI.UpdateServices(ctx2, service.ConvertServiceStepUpdateAgentToManager(up)); err != nil {
		log.Errorf("SynapseAgentPod Rebounce: failed to update service status(success): error: %s\n", err.Error())
	}

	f.Cancel()
}
