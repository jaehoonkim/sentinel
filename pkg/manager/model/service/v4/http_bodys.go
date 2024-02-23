package service

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	v3 "github.com/jaehoonkim/morpheus/pkg/manager/model/service/v3"
)

// HttpRsp_AgentServicePolling
//
//	http responce body; agent service polling
type HttpRsp_AgentServicePolling_multistep = Service

// HttpReq_AgentServiceUpdate_multistep
//
//	http request body; agent service update
type HttpReq_AgentServiceUpdate_multistep struct {
	Uuid     string     `json:"uuid"`     //pk
	Sequence int        `json:"sequence"` //pk
	Status   StepStatus `json:"status"`
	Result   string     `json:"result"` //StepStatus 값에 따라; 결과 혹은 에러 메시지
	Started  time.Time  `json:"started"`
	Ended    time.Time  `json:"ended"`
}

type HttpRsp_Service struct {
	Service       `json:",inline"`
	StatusRecords []ServiceStatus `json:"status_records,omitempty"`
}

type HttpReq_Service_create struct {
	ClusterUuid       []string               `json:"cluster_uuid,omitempty"`
	Name              string                 `json:"name,omitempty"`
	Summary           string                 `json:"summary,omitempty"`
	TemplateUuid      string                 `json:"template_uuid,omitempty"`
	Inputs            map[string]interface{} `json:"inputs,omitempty"`
	SubscribedChannel string                 `json:"subscribed_channel,omitempty"`

	IsMultiCluster bool `json:"-"`
	IsMultiSteps   bool `json:"-"`
}

type HttpRsp_Service_create = Service

func (obj HttpReq_Service_create) MarshalJSON() ([]byte, error) {
	return json.Marshal(obj)
}

func (obj *HttpReq_Service_create) UnmarshalJSON(bytes []byte) error {

	type T struct {
		ClusterUuid       json.RawMessage        `json:"cluster_uuid,omitempty"`
		Name              string                 `json:"name,omitempty"`
		Summary           string                 `json:"summary,omitempty"`
		TemplateUuid      string                 `json:"template_uuid,omitempty"`
		Inputs            map[string]interface{} `json:"inputs,omitempty"`
		SubscribedChannel string                 `json:"subscribed_channel,omitempty"`
		Steps             []struct {
			Args map[string]interface{} `json:"args,omitempty"`
		} `json:"steps,omitempty"`
	}

	var v T
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}

	var cluster_uuid = []string{}
	if 0 < len(v.ClusterUuid) &&
		string(v.ClusterUuid)[0] == '[' &&
		string(v.ClusterUuid)[len(string(v.ClusterUuid))-1] == ']' {

		if err := json.Unmarshal(v.ClusterUuid, &cluster_uuid); err != nil {
			return err
		}

		obj.IsMultiCluster = true
	}

	if 0 < len(v.ClusterUuid) &&
		string(v.ClusterUuid)[0] == '"' &&
		string(v.ClusterUuid)[len(string(v.ClusterUuid))-1] == '"' {

		s := string(v.ClusterUuid)
		s, _ = strconv.Unquote(s)
		s = strings.TrimSpace(s)

		cluster_uuid = append(cluster_uuid, s)

		obj.IsMultiCluster = false
	}

	obj.ClusterUuid = cluster_uuid
	obj.Name = v.Name
	obj.Summary = v.Summary
	obj.TemplateUuid = v.TemplateUuid
	obj.Inputs = v.Inputs
	if 0 < len(v.Steps) {
		obj.Inputs = v.Steps[0].Args
	}

	obj.SubscribedChannel = v.SubscribedChannel

	obj.IsMultiSteps = !(0 < len(v.Steps))

	return nil
}

type HttpRsp_ServiceStatues = ServiceStatus

type HttpRsp_ServiceResult = ServiceResult

type HttpRsp_AgentServicePolling struct {
	Version string `json:"version" enum:"v3,v4"`

	V3 v3.HttpRsp_AgentServicePolling        `json:"-"`
	V4 HttpRsp_AgentServicePolling_multistep `json:"-"`
}

func (obj HttpRsp_AgentServicePolling) MarshalJSON() ([]byte, error) {

	type T struct {
		Version string          `json:"version" enum:"v3,v4"`
		Service json.RawMessage `json:"service,omitempty"`
	}

	var v T
	v.Version = obj.Version

	switch obj.Version {
	case "v3":
		b, err := json.Marshal(obj.V3)
		if err != nil {
			return nil, err
		}
		v.Service = b
	case "v4":
		b, err := json.Marshal(obj.V4)
		if err != nil {
			return nil, err
		}
		v.Service = b
	default:
		err := errors.New("invalid version")
		return nil, err
	}

	return json.Marshal(v)
}

func (obj *HttpRsp_AgentServicePolling) UnmarshalJSON(bytes []byte) error {
	type T struct {
		Version string          `json:"version" enum:"v3,v4"`
		Service json.RawMessage `json:"service,omitempty"`
	}

	var v T
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}

	obj.Version = v.Version

	switch v.Version {
	case "v3":
		if err := json.Unmarshal(v.Service, &obj.V3); err != nil {
			return err
		}
	case "v4":
		if err := json.Unmarshal(v.Service, &obj.V4); err != nil {
			return err
		}
	default:
		err := errors.New("invalid version")
		return err
	}

	return nil
}

type HttpReq_AgentServiceUpdate struct {
	Version string `json:"version" enum:"v3,v4"`

	V3 v3.HttpReq_AgentServiceUpdate        `json:"-"`
	V4 HttpReq_AgentServiceUpdate_multistep `json:"-"`
}

func (obj HttpReq_AgentServiceUpdate) MarshalJSON() ([]byte, error) {

	type T struct {
		Version string          `json:"version"`
		Service json.RawMessage `json:"service,omitempty"`
	}

	var v T
	v.Version = obj.Version

	switch obj.Version {
	case "v3":
		b, err := json.Marshal(obj.V3)
		if err != nil {
			return nil, err
		}
		v.Service = b
	case "v4":
		fallthrough
	default:
		b, err := json.Marshal(obj.V4)
		if err != nil {
			return nil, err
		}
		v.Service = b
	}

	return json.Marshal(v)
}

func (obj *HttpReq_AgentServiceUpdate) UnmarshalJSON(bytes []byte) error {
	type T struct {
		Version string          `json:"version"`
		Service json.RawMessage `json:"service,omitempty"`
	}

	var v T
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}

	obj.Version = v.Version

	switch v.Version {
	case "v3":
		if err := json.Unmarshal(v.Service, &obj.V3); err != nil {
			return err
		}
	case "v4":
		fallthrough
	default:
		if err := json.Unmarshal(v.Service, &obj.V4); err != nil {
			return err
		}
	}

	return nil
}
