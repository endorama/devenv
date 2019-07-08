package profile

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

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
	return "email"
}

func (p EmailPlugin) Render(profile Profile) string {
	config := p.Config().(EmailPluginConfig)
	sb := strings.Builder{}
	sb.WriteString("export EMAIL=" + config.Email + "\n")
	return sb.String()
}

func (p EmailPlugin) Config() interface{} {
	return p.config
}

func (p EmailPlugin) ConfigFile(profileLocation string) string {
	return path.Join(profileLocation, "config-"+p.Name()+".yaml")
}

func (p *EmailPlugin) LoadConfig(profileLocation string) error {
	content, err := ioutil.ReadFile(p.ConfigFile(profileLocation))
	if err != nil {
		return errors.Wrap(err, "cannot read config file")
	}
	err = yaml.Unmarshal([]byte(content), &p.config)
	if err != nil {
		return errors.Wrap(err, "cannot unmarshal config file")
	}
	return nil
}
