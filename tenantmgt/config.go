package usermgt

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port string `yaml:"port"`
}

type MongoConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	ServerConfig ServerConfig `yaml:"server"`
}

func ReadConfig() (Config, error) {
	var config Config
	var err error
	configBytes, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Println("unable to read config file: ", err)
		return Config{}, err
	}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		log.Println("unable to unmarshal config: ", err)
		return Config{}, err
	}
	return config, nil
}

func GetPort() (string, error) {
	config, err := ReadConfig()
	if err != nil {
		return "", err
	}
	return config.ServerConfig.Port, nil
}
