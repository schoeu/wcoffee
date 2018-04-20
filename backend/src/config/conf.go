package config

import (
	"fmt"
	"github.com/schoeu/golodash"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Conf struct {
	Port string `yaml:"port"`
	DBString string `yaml:"dbstring"`
}

var globalConfig *Conf

func init() {
	var conf Conf
	globalConfig = conf.GetConfig()
}

func (c *Conf) GetConfig() *Conf {
	confFileName := "config.yaml"
	projectPath := filepath.Join(golodash.GetCwd(), ".", confFileName)
	yamlFile, err := ioutil.ReadFile(projectPath)

	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}

	return c
}

func GetConf() *Conf {
	return globalConfig
}
