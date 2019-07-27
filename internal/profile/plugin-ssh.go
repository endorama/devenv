package profile

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	pluginSSHDefaultCachePath = "/tmp/%s-ssh-agent.tmp"
)

type SSHPlugin struct {
	config SSHPluginConfig
}

type SSHPluginConfig struct {
	Keys      []string
	CachePath string
}

func NewSSHPlugin() *SSHPlugin {
	return &SSHPlugin{}
}

func (p SSHPlugin) Name() string {
	return "ssh"
}

func (p SSHPlugin) Render(profile Profile) string {
	config := p.Config().(SSHPluginConfig)

	sshPluginFolder := path.Join(profile.Location, p.Name())

	if config.CachePath == "" {
		config.CachePath = pluginSSHDefaultCachePath
	}
	// cache should be by profile
	config.CachePath = fmt.Sprintf(config.CachePath, profile.Name)

	// to avoid specifying the entire path to the key, we expect them to be in
	// {{profileLocation}}/ssh
	for idx, value := range config.Keys {
		config.Keys[idx] = path.Join(sshPluginFolder, value)
	}

	t := `# create agent cache if missing
if [ ! -f {{.CachePath}} ]; then
	echo $(ssh-agent -s | sed "s/echo/# echo/") > {{.CachePath}}
	chown $USER:$USER {{.CachePath}}
	chmod 600 {{.CachePath}}
fi
# load agent
source {{.CachePath}}
# add ssh keys, if not already loaded
{{ range $key, $value := .Keys -}}
if ! ssh-add -l 2> /dev/null | grep {{$value}} > /dev/null; then
	ssh-add {{$value}} > /dev/null
fi
{{end -}}`

	sb := strings.Builder{}
	tpl, err := template.New("ssh-plugins").Parse(t)
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(&sb, config)

	return sb.String()
}

func (p SSHPlugin) Config() interface{} {
	return p.config
}

func (p SSHPlugin) ConfigFile(profileLocation string) string {
	return path.Join(profileLocation, "config-"+p.Name()+".yaml")
}

func (p *SSHPlugin) LoadConfig(profileLocation string) error {
	content, err := ioutil.ReadFile(p.ConfigFile(profileLocation))
	if err != nil {
		return errors.Wrap(err, "(ssh) cannot read config file")
	}
	err = yaml.Unmarshal([]byte(content), &p.config)
	if err != nil {
		return errors.Wrap(err, "(ssh) cannot unmarshal config file")
	}
	return nil
}
