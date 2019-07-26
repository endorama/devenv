package profile

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type EnvsPlugin struct {
	// Pluggable, Configurable
	config EnvsPluginConfig
}

type EnvsPluginConfig struct {
	Envs map[string]string
}

func NewEnvsPlugin() *EnvsPlugin {
	return &EnvsPlugin{}
}

func (p EnvsPlugin) Name() string {
	return "envs"
}

func (p EnvsPlugin) Render(profile Profile) string {
	config := p.Config().(EnvsPluginConfig)
	sb := strings.Builder{}
	for name, value := range config.Envs {
		name = strings.ToUpper(name)
		sb.WriteString(fmt.Sprintf("export %s=\"%s\"\n", name, value))
	}
	return sb.String()
}

func (p EnvsPlugin) Config() interface{} {
	return p.config
}

func (p EnvsPlugin) ConfigFile(profileLocation string) string {
	return path.Join(profileLocation, "config-"+p.Name()+".yaml")
}

func (p *EnvsPlugin) LoadConfig(profileLocation string) error {
	content, err := ioutil.ReadFile(p.ConfigFile(profileLocation))
	if err != nil {
		return errors.Wrap(err, "(envs) cannot read config file")
	}
	err = yaml.Unmarshal([]byte(content), &p.config)
	if err != nil {
		return errors.Wrap(err, "(envs) cannot unmarshal config file")
	}
	return nil
}
