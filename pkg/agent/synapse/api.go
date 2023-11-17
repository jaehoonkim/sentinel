package synapse

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/jaehoonkim/synapse/pkg/agent/httpclient"
	"github.com/jaehoonkim/synapse/pkg/agent/log"
	"github.com/jaehoonkim/synapse/pkg/manager/model/auths/v2"
	servicev3 "github.com/jaehoonkim/synapse/pkg/manager/model/service/v3"
	servicev4 "github.com/jaehoonkim/synapse/pkg/manager/model/service/v4"
	sessionv1 "github.com/jaehoonkim/synapse/pkg/manager/model/session/v1"
)

const synapseAuthTokenHeaderName = "x-synapse-agent-token"

type SynapseAPI struct {
	client    *httpclient.HttpClient
	authToken atomic.Value
}

func NewSynapseAPI(address string) (*SynapseAPI, error) {
	log.Debugf("address in NewSynapseAPI - api.go : %s\n", address)
	client, err := httpclient.NewHttpClient(address, false, 0, 0)
	if err != nil {
		return nil, err
	}
	log.Debugf("client in NewSynapseAPI - api.go : %s\n", client)
	return &SynapseAPI{client: client}, nil
}

func NewSynapseAPIWithClient(client *httpclient.HttpClient) *SynapseAPI {
	return &SynapseAPI{client: client}
}

func (s *SynapseAPI) IsTokenExpired() bool {
	claims := new(sessionv1.ClientSessionPayload)
	jwt_token, _, err := jwt.NewParser().ParseUnverified(s.GetToken(), claims)
	if _, ok := jwt_token.Claims.(*sessionv1.ClientSessionPayload); !ok || err != nil {
		log.Warnf("jwt.ParseUnverified error : %v\n", err)
		return true
	}

	return !claims.VerifyExpiresAt(time.Now().Unix(), true)
}

func (s *SynapseAPI) GetToken() string {
	x := s.authToken.Load()

	return x.(string)
}

func (s *SynapseAPI) Auth(ctx context.Context, auth *auths.HttpReqAuth) error {
	if auth == nil {
		return fmt.Errorf("auth is nil")
	}

	b, err := json.Marshal(auth)
	if err != nil {
		return err
	}
	log.Debugf("synapse API in Auth : %s\n", s.client)
	result := s.client.Post("/client/auth").SetBody("application/json", b).Do(ctx)

	// get session token
	if headers := result.Headers(); headers != nil {
		if token := headers.Get(synapseAuthTokenHeaderName); token != "" {
			s.authToken.Store(token)
		}
	}

	if err := result.Error(); err != nil {
		return err
	}

	return nil
}

func (s *SynapseAPI) GetServices(ctx context.Context) ([]servicev4.HttpRsp_AgentServicePolling, error) {
	var services HttpPollingDataset

	token := s.GetToken()
	if token == "" {
		return nil, fmt.Errorf("session token is empty")
	}

	result := s.client.Get("/client/service").
		SetHeader(synapseAuthTokenHeaderName, token).
		Do(ctx)

	// get session token
	if headers := result.Headers(); headers != nil {
		if token := headers.Get(synapseAuthTokenHeaderName); token != "" {
			s.authToken.Store(token)
		}
	}

	if err := result.IntoJson(&services); err != nil {
		return nil, err
	}

	return services, nil
}

func (s *SynapseAPI) UpdateServices(ctx context.Context, service *servicev4.HttpReq_AgentServiceUpdate) error {
	if service == nil {
		return fmt.Errorf("service is nil")
	}

	switch service.Version {
	case "v3":
		log.Debugf("request update_service: service{version:%s, uuid:%s, status:%d, result_len:%d}\n", service.Version, service.V3.Uuid, service.V3.Status, len(service.V3.Result))
	case "v4":
		log.Debugf("request update_service: service{version:%s, uuid:%s, status:%d, result_len:%d}\n", service.Version, service.V4.Uuid, service.V4.Status, len(service.V4.Result))
	default:
		return fmt.Errorf("unknown service version: %s", service.Version)
	}

	b, err := json.Marshal(service)
	if err != nil {
		return err
	}

	token := s.GetToken()
	if token == "" {
		return fmt.Errorf("session token is empty")
	}

	result := s.client.Put("/client/service").
		SetHeader(synapseAuthTokenHeaderName, token).
		SetGzip(true).
		SetBody("application/json", b).
		Do(ctx)

	// get session token
	if headers := result.Headers(); headers != nil {
		if token := headers.Get(synapseAuthTokenHeaderName); token != "" {
			s.authToken.Store(token)
		}
	}

	if err := result.Error(); err != nil {
		return err
	}

	return nil
}

type HttpPollingDataset []servicev4.HttpRsp_AgentServicePolling

func (s *HttpPollingDataset) UnmarshalJSON(b []byte) error {
	var l []json.RawMessage

	if err := json.Unmarshal(b, &l); err != nil {
		return err
	}

	for _, e := range l {
		data := servicev4.HttpRsp_AgentServicePolling{}
		if err := json.Unmarshal(e, &data); err != nil {
			// older version
			datav3 := servicev3.HttpRsp_AgentServicePolling{}
			if err := json.Unmarshal(e, &datav3); err != nil {
				return err
			}
			data.Version = "v3"
			data.V3 = datav3
		}

		*s = append(*s, data)
	}

	return nil
}
