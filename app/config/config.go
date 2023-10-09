package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type ConfigType struct {
	Port           uint          `yaml:"port"`
	Redis          RedisType     `yaml:"redis"`
	MySQL          MySQLType     `yaml:"mysql"`
	Dcron          DcronType     `yaml:"dcron"`
	InnerCall      InnerCallType `yaml:"innerCall"`
	EnableReporter bool          `yaml:"enableReporter"`
}

type RedisType struct {
	Addr     string `yaml:"addr"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type MySQLType struct {
	User     string `yaml:"user"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

type DcronType struct {
	ServiceName string `yaml:"serviceName"`
}

type InnerCallType struct {
	Channel string    `yaml:"channel"`
	Redis   RedisType `yaml:"redis"`
}

var (
	configInstance ConfigType
)

func I() ConfigType {
	return configInstance
}

func LoadConfig(filename string) (err error) {
	return loadConfigFromFile(filename, &configInstance)
}

func loadConfigFromFile(filename string, config *ConfigType) (err error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, config)
	return
}
