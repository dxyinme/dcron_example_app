package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type ConfigType struct {
	Port  uint      `yaml:"port"`
	Redis RedisType `yaml:"redis"`
	MySQL MySQLType `yaml:"mysql"`
	Rpc   RpcType   `yaml:"rpc"`
}

type RedisType struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

type MySQLType struct {
	User     string `yaml:"user"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

type RpcType struct {
	Port uint `yaml:"port"`
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