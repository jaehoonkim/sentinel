package executor

import (
	"context"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/gophercloud/utils/openstack/clientconfig"
	"github.com/jaehoonkim/morpheus/pkg/agent/alertmanager"
	"github.com/jaehoonkim/morpheus/pkg/agent/grafana"
	"github.com/jaehoonkim/morpheus/pkg/agent/helm"
	"github.com/jaehoonkim/morpheus/pkg/agent/jq"
	"github.com/jaehoonkim/morpheus/pkg/agent/k8s"
	"github.com/jaehoonkim/morpheus/pkg/agent/morpheusagent"
	"github.com/jaehoonkim/morpheus/pkg/agent/openstack"
	"github.com/jaehoonkim/morpheus/pkg/agent/p8s"
	"github.com/jaehoonkim/morpheus/pkg/agent/service"
	"github.com/jaehoonkim/morpheus/pkg/manager/macro"
)

type CommandType int

const (
	CommandTypeK8s CommandType = iota + 1
	CommandTypeP8s
	CommandTypeHelm
	CommandTypeJq
	CommandTypeAlertManager
	CommandTypeSynapseagent
	CommandTypeGrafana
	CommandTypeOpenstack
)

func (ct CommandType) String() string {
	if ct == CommandTypeK8s {
		return "kubernetes"
	} else if ct == CommandTypeP8s {
		return "prometheus"
	} else if ct == CommandTypeHelm {
		return "helm"
	} else if ct == CommandTypeJq {
		return "jq"
	} else if ct == CommandTypeAlertManager {
		return "alertmanager"
	} else if ct == CommandTypeSynapseagent {
		return "morpheus"
	} else if ct == CommandTypeGrafana {
		return "grafana"
	} else if ct == CommandTypeOpenstack {
		return "openstack"
	}

	return "Unknown CommandType"
}

type Commander interface {
	GetCommandType() CommandType
	Run() (string, error)
}

func NewCommander(command *service.StepCommand) (Commander, error) {
	mlist := strings.Split(command.Method, ".")
	ctype := mlist[0]

	switch ctype {
	case "kubernetes":
		return NewK8sCommander(command)
	case "prometheus":
		return NewP8sCommander(command)
	case "helm":
		return NewHelmCommander(command)
	case "jq":
		return NewJqCommander(command)
	case "alertmanager":
		return NewAlertManagerCommander(command)
	case "morpheus":
		return NewSynapseagentCommander(command)
	case "grafana":
		return NewGrafanaCommander(command)
	case "openstack":
		return NewOpenstackCommander(command)
	}

	return nil, fmt.Errorf("unknown command method(%s)", command.Method)
}

type K8sCommander struct {
	client   *k8s.Client
	gv       schema.GroupVersion // v1, apps/v1, ...
	resource string              // pod, namespace, deployment, ...
	verb     string              // get, list, watch, ...
	args     map[string]interface{}
}

func NewK8sCommander(command *service.StepCommand) (Commander, error) {
	client, err := k8s.GetClient()
	if err != nil {
		return nil, err
	}
	cmdr := &K8sCommander{client: client, args: make(map[string]interface{})}

	err = cmdr.ParseCommand(command)
	if err != nil {
		return nil, err
	}

	return cmdr, nil
}

func (c *K8sCommander) GetCommandType() CommandType {
	return CommandTypeK8s
}

func (c *K8sCommander) ParseCommand(command *service.StepCommand) error {
	mlist := strings.SplitN(command.Method, ".", 4)

	if len(mlist) != 4 {
		return fmt.Errorf("there is not enough method(%s) for k8s. want(4) but got(%d)", command.Method, len(mlist))
	}

	gv, err := schema.ParseGroupVersion(mlist[3])
	if err != nil {
		return err
	}

	c.gv = gv
	c.resource = mlist[1]
	c.verb = mlist[2]

	if command.Args != nil {
		c.args = command.Args
	}

	return nil
}

func (c *K8sCommander) Run() (string, error) {
	return c.client.ResourceRequest(c.gv, c.resource, c.verb, c.args)
}

type P8sCommander struct {
	client      *p8s.Client
	apiVersion  string
	api         string
	queryParams map[string]interface{}
}

func NewP8sCommander(command *service.StepCommand) (Commander, error) {
	cmdr := &P8sCommander{}

	if err := cmdr.ParseCommand(command); err != nil {
		return nil, err
	}

	return cmdr, nil
}

func (c *P8sCommander) GetCommandType() CommandType {
	return CommandTypeP8s
}

func (c *P8sCommander) ParseCommand(command *service.StepCommand) error {
	mlist := strings.SplitN(command.Method, ".", 3)

	if len(mlist) != 3 {
		return fmt.Errorf("there is not enough method(%s) for p8s. want(3) but got(%d)", command.Method, len(mlist))
	}

	c.api = mlist[1]
	c.apiVersion = mlist[2]
	c.queryParams = command.Args

	url, ok := macro.MapString(command.Args, "url")
	if !ok || len(url) == 0 {
		return fmt.Errorf("prometheus url is empty")
	}

	client, err := p8s.NewClient(url)
	if err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *P8sCommander) Run() (string, error) {
	return c.client.ApiRequest(c.apiVersion, c.api, c.queryParams)
}

type HelmCommander struct {
	client *helm.Client
	cmd    string
	args   map[string]interface{}
}

func NewHelmCommander(command *service.StepCommand) (Commander, error) {
	cmdr := &HelmCommander{}

	if err := cmdr.ParseCommand(command); err != nil {
		return nil, err
	}

	return cmdr, nil
}

func (c *HelmCommander) GetCommandType() CommandType {
	return CommandTypeHelm
}

func (c *HelmCommander) ParseCommand(command *service.StepCommand) error {
	mlist := strings.SplitN(command.Method, ".", 2)

	if len(mlist) != 2 {
		return fmt.Errorf("there is not enough method(%s) for helm. want(3) but got(%d)", command.Method, len(mlist))
	}

	c.cmd = mlist[1]
	c.args = command.Args

	client, err := helm.NewClient()
	if err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *HelmCommander) Run() (string, error) {
	return c.client.Request(c.cmd, c.args)
}

type JqCommander struct {
	input  interface{}
	filter string
}

func NewJqCommander(command *service.StepCommand) (Commander, error) {
	cmdr := &JqCommander{}

	if err := cmdr.ParseCommand(command); err != nil {
		return nil, err
	}

	return cmdr, nil
}

func (c *JqCommander) GetCommandType() CommandType {
	return CommandTypeJq
}

func (c *JqCommander) ParseCommand(command *service.StepCommand) error {
	if m, ok := command.Args["input"]; ok {
		c.input = m
	} else {
		return fmt.Errorf("input not found")
	}

	if f, ok := macro.MapString(command.Args, "filter"); ok {
		c.filter = f
	} else {
		return fmt.Errorf("filter not found")
	}

	return nil
}

func (c *JqCommander) Run() (string, error) {
	return jq.Request(c.input, c.filter)
}

type AlertManagerCommander struct {
	client     *alertmanager.Client
	apiVersion string
	api        string
	verb       string
	params     map[string]interface{}
}

func NewAlertManagerCommander(command *service.StepCommand) (Commander, error) {
	cmdr := &AlertManagerCommander{}

	if err := cmdr.ParseCommand(command); err != nil {
		return nil, err
	}

	return cmdr, nil
}

func (c *AlertManagerCommander) GetCommandType() CommandType {
	return CommandTypeAlertManager
}

func (c *AlertManagerCommander) ParseCommand(command *service.StepCommand) error {
	mlist := strings.SplitN(command.Method, ".", 4)

	if len(mlist) != 4 {
		return fmt.Errorf("there is not enough method(%s) for alertmanager. want(4) but got(%d)", command.Method, len(mlist))
	}

	c.api = mlist[1]
	c.verb = mlist[2]
	c.apiVersion = mlist[3]
	c.params = command.Args

	url, ok := macro.MapString(command.Args, "url")
	if !ok || len(url) == 0 {
		return fmt.Errorf("alertmanager url is empty")
	}

	client, err := alertmanager.NewClient(url)
	if err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *AlertManagerCommander) Run() (string, error) {
	return c.client.ApiRequest(c.apiVersion, c.api, c.verb, c.params)
}

type SynapseagentCommander struct {
	client *morpheusagent.Client
	api    string
	verb   string
	params map[string]interface{}
}

func NewSynapseagentCommander(command *service.StepCommand) (Commander, error) {
	cmdr := &SynapseagentCommander{}

	if err := cmdr.ParseCommand(command); err != nil {
		return nil, err
	}

	return cmdr, nil
}

func (c *SynapseagentCommander) GetCommandType() CommandType {
	return CommandTypeSynapseagent
}

func (c *SynapseagentCommander) ParseCommand(command *service.StepCommand) error {
	mlist := strings.SplitN(command.Method, ".", 3)

	if len(mlist) != 3 {
		return fmt.Errorf("there is not enough method(%s) for morpheusagent. want(3) but got(%d)", command.Method, len(mlist))
	}

	c.api = mlist[1]
	c.verb = mlist[2]
	c.params = command.Args

	client, err := morpheusagent.NewClient()
	if err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *SynapseagentCommander) Run() (string, error) {
	return c.client.Request(c.api, c.verb, c.params)
}

type GrafanaCommander struct {
	client *grafana.Client
	api    string
	verb   string
	params map[string]interface{}
}

func NewGrafanaCommander(command *service.StepCommand) (Commander, error) {
	cmdr := &GrafanaCommander{}

	if err := cmdr.ParseCommand(command); err != nil {
		return nil, err
	}

	return cmdr, nil
}

func (c *GrafanaCommander) GetCommandType() CommandType {
	return CommandTypeGrafana
}

func (c *GrafanaCommander) ParseCommand(command *service.StepCommand) error {
	mlist := strings.SplitN(command.Method, ".", 3)

	if len(mlist) != 3 {
		return fmt.Errorf("there is not enough method(%s) for grafana. want(3) but got(%d)", command.Method, len(mlist))
	}

	c.api = mlist[1]
	c.verb = mlist[2]
	c.params = command.Args

	url, ok := macro.MapString(command.Args, "url")
	if !ok || len(url) == 0 {
		return fmt.Errorf("grafana url is empty")
	}

	credentialKey, ok := macro.MapString(command.Args, "credential_key")
	if !ok || len(credentialKey) == 0 {
		return fmt.Errorf("grafana credential_key is empty")
	}

	client, err := grafana.NewClient(url, func() ([]byte, error) {
		kc, err := k8s.GetClient()
		if err != nil {
			return nil, err
		}

		secret, err := kc.GetK8sClientset().CoreV1().Secrets("morpheusagent").Get(context.Background(), morpheusagent.SynapseagentSecretName, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}

		credentialYaml, ok := secret.Data[credentialKey]
		if !ok || len(credentialYaml) <= 0 {
			return nil, fmt.Errorf("could not find apikey from credential_key(%s)", credentialKey)
		}

		type GrafanaCredential struct {
			Type string `yaml:"type"`
			Data string `yaml:"data"`
		}

		gc := new(GrafanaCredential)
		if err := yaml.Unmarshal(credentialYaml, gc); err != nil {
			return nil, err
		}

		return []byte(gc.Data), nil
	})
	if err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *GrafanaCommander) Run() (string, error) {
	return c.client.ApiRequest(c.api, c.verb, c.params)
}

type OpenstackCommander struct {
	client   *openstack.Client
	api      string
	resource string
	verb     string
	params   map[string]interface{}
}

func NewOpenstackCommander(command *service.StepCommand) (Commander, error) {
	cmdr := &OpenstackCommander{}

	if err := cmdr.ParseCommand(command); err != nil {
		return nil, err
	}

	return cmdr, nil
}

func (c *OpenstackCommander) GetCommandType() CommandType {
	return CommandTypeOpenstack
}

func (c *OpenstackCommander) ParseCommand(command *service.StepCommand) error {
	mlist := strings.SplitN(command.Method, ".", 4)

	if len(mlist) != 4 {
		return fmt.Errorf("there is not enough method(%s) for openstack. want(4) but got(%d)", command.Method, len(mlist))
	}

	c.api = mlist[1]
	c.resource = mlist[2]
	c.verb = mlist[3]

	c.params = command.Args

	credentialKey, ok := macro.MapString(command.Args, "credential_key")
	if !ok || len(credentialKey) == 0 {
		return fmt.Errorf("openstack credential_key is empty")
	}

	kc, err := k8s.GetClient()
	if err != nil {
		return err
	}

	secret, err := kc.GetK8sClientset().CoreV1().Secrets(morpheusagent.SynapseagentNamespace).Get(context.Background(), morpheusagent.SynapseagentSecretName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	credentialYaml, ok := secret.Data[credentialKey]
	if !ok || len(credentialYaml) <= 0 {
		return fmt.Errorf("could not find apikey from credential_key(%s)", credentialKey)
	}

	type OpenstackCredential struct {
		Type string               `yaml:"type"`
		Data *clientconfig.Clouds `yaml:"data"`
	}

	oc := new(OpenstackCredential)
	if err := yaml.Unmarshal(credentialYaml, oc); err != nil {
		return err
	}

	if oc.Data == nil {
		return fmt.Errorf("openstack data is nil")
	}

	for _, cloud := range oc.Data.Clouds {
		opts := &clientconfig.ClientOpts{
			AuthInfo: cloud.AuthInfo,
		}

		pClient, err := clientconfig.AuthenticatedClient(opts)
		if err != nil {
			return err
		}

		c.client = openstack.NewClient(pClient)

		return nil
	}

	return fmt.Errorf("openstack clouds.yaml auth_info is empty")
}

func (c *OpenstackCommander) Run() (string, error) {
	return c.client.ApiRequest(c.api, c.resource, c.verb, c.params)
}
