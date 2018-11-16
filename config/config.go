package config

import (
	"io/ioutil"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

type Config struct {
	MQNameSvcAddr string `yaml:"maNameSvcAddr"`
	MQTopic       string `yaml:"mqTopic"`
	DB            string `yaml:"db"`
	MaxRows       int32  `yaml:"maxRows"`
}

var cfg *Config

func ParseConfig(configFile string) *Config {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		glog.Exitf("Failed to read config file %s: %v", configFile, err)
	}

	var config *Config
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		glog.Exitf("Failed to unmarshal config %v", err)
	}

	return config
}

func GetConfig() *Config {
	return cfg
}
