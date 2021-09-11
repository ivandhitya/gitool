package example

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Gitlab struct {
		Address string `yaml:"address"`
		Token   string `yaml:"token"`
	} `yaml:"gitlab"`
}

func (c *Conf) GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile(".gitool.yaml")
	if err != nil {
		logrus.Error("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		logrus.Panic("Unmarshal: %v", err)
	}
	return c
}
