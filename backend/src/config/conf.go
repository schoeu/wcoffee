package config

import (
	"../utils"
	"github.com/schoeu/golodash"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Conf struct {
	Port     string `yaml:"port"`
	DBString string `yaml:"dbstring"`
	Mode     string `yaml:"mode"`
}

var globalConfig *Conf

func init() {
	var conf Conf
	globalConfig = conf.GetConfig()
}

func (c *Conf) GetConfig() *Conf {
	confFileName := "config.yaml"
	projectPath := filepath.Join(golodash.GetCwd(), "backend", confFileName)
	yamlFile, err := ioutil.ReadFile(projectPath)
	utils.ErrHandle(err)

	err = yaml.Unmarshal(yamlFile, c)
	utils.ErrHandle(err)

	return c
}

func GetConf() *Conf {
	return globalConfig
}
