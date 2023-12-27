package config

import (
	"fmt"
	"os"
	"strings"

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

// channel name and channel redis configuration,
// TODO:
// remove redis dependency in inner call.
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

func LoadConfig(mode, filename string) (err error) {
	switch strings.ToLower(mode) {
	case "fromfile":
		return loadConfigFromFile(filename, &configInstance)
	case "fromenv":
		return loadConfigFromEnv(&configInstance)
	default:
		return fmt.Errorf("mode error: %s", mode)
	}
}

func loadConfigFromFile(filename string, config *ConfigType) (err error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, config)
	return
}

func loadConfigFromEnv(config *ConfigType) (err error) {
	return
}
