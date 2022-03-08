package helm

import (
	"encoding/json"
	"fmt"
	"time"

	"helm.sh/helm/v3/pkg/action"
)

func setDefaultUninstallSettings(client *action.Uninstall) {
	client.Wait = true
	client.Timeout = 300 * time.Second
}

func (c *Client) Uninstall(args map[string]interface{}) (string, error) {
	type UninstallParams struct {
		Namespace string `param:"namespace"`
		Name      string `param:"name"`
	}

	params := &UninstallParams{}

	if err := convertArgsToStruct(args, params); err != nil {
		return "", err
	}

	// set namespace
	c.settings.SetNamespace(params.Namespace)

	// get uninstall action client
	actionConfig, err := c.getActionConfig()
	if err != nil {
		return "", err
	}
	client := action.NewUninstall(actionConfig)

	// default settings
	setDefaultUninstallSettings(client)

	resp, err := client.Run(params.Name)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return fmt.Sprintf("chart(%s) uninstall is success, but failed to json.Marhsal : %s", params.Name, err.Error()), nil
	}

	return string(b), nil
}
