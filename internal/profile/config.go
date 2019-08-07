package profile

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type ProfileConfig struct {
	Plugins []string
}

func (pc *ProfileConfig) LoadFromFile(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "cannot read yaml file")
	}
	err = yaml.Unmarshal(yamlFile, pc)
	if err != nil {
		return errors.Wrap(err, "cannot unmarshal yaml file")
	}
	return nil
}
