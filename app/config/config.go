package config

import (
	"fmt"
	"os"
	"strconv"
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

func fromEnvString(envStr string, defaultValue string) string {
	s := os.Getenv(envStr)
	if len(s) == 0 {
		return defaultValue
	}
	return s
}

func fromEnvInt(envStr string, defaultValue int) int {
	s := os.Getenv(envStr)
	if len(s) == 0 {
		return defaultValue
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func loadConfigFromEnv(config *ConfigType) (err error) {

	config.Dcron.ServiceName = fromEnvString("APP_DCRON_SERVICENAME", "dcronapp")
	{
		EnableReporterStr := fromEnvString("APP_ENABLEREPORTER", "false")
		if strings.ToLower(EnableReporterStr) == "true" {
			config.EnableReporter = true
		} else {
			config.EnableReporter = false
		}
	}

	config.InnerCall.Channel = fromEnvString("APP_INNERCALL_CHANNEL", "dcronapp-channel")
	config.InnerCall.Redis.Addr = fromEnvString("APP_INNERCALL_REDIS_ADDR", "localhost:2379")
	config.InnerCall.Redis.DB = fromEnvInt("APP_INNERCALL_REDIS_DB", 2)
	config.InnerCall.Redis.Password = fromEnvString("APP_INNERCALL_REDIS_PASSWORD", "")

	config.MySQL.Addr = fromEnvString("APP_MYSQL_ADDR", "")
	config.MySQL.Password = fromEnvString("APP_MYSQL_PASSWORD", "")
	config.MySQL.User = fromEnvString("APP_MYSQL_USER", "")

	config.Port = uint(fromEnvInt("APP_PORT", 8080))

	config.Redis.Addr = fromEnvString("APP_REDIS_ADDR", "localhost:2379")
	config.Redis.DB = fromEnvInt("APP_REDIS_DB", 0)
	config.Redis.Password = fromEnvString("APP_REDIS_PASSWORD", "")

	return
}
