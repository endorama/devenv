package profile

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const EmailPluginName = "email"

type EmailPlugin struct {
	config EmailPluginConfig
}

type EmailPluginConfig struct {
	Email string `yaml:"email"`
}

func NewEmailPlugin() *EmailPlugin {
	return &EmailPlugin{}
}

func (p EmailPlugin) Name() string {
	return EmailPluginName
}

func (p EmailPlugin) Render(profileName, profileLocation string) string {
	config := p.Config().(EmailPluginConfig)
	sb := strings.Builder{}
	sb.WriteString("export EMAIL=" + config.Email + "\n")
	return sb.String()
}

func (p EmailPlugin) Config() interface{} {
	return p.config
}

func (p EmailPlugin) ConfigFile(profileLocation string) string {
	return path.Join(profileLocation, "config-"+EmailPluginName+".yaml")
}

func (p *EmailPlugin) LoadConfig(profileLocation string) error {
	configFile := p.ConfigFile(profileLocation)
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return errors.Wrap(err, "cannot read config file")
	}
	err = yaml.Unmarshal([]byte(content), &p.config)
	if err != nil {
		return errors.Wrap(err, "cannot unmarshal config file")
	}
	return nil
}
