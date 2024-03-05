package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	configReader = viper.New()
)

type Config struct {
	Server Server `mapstructure:"server"`
	Mongo  Mongo  `mapstructure:"mongo"`
}

type Mongo struct {
	Uri      string `mapstructure:"uri"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Server struct {
	Port         string `mapstructure:"port"`
	Issuer       string `mapstructure:"issuer"`
	FirstSeed    bool   `mapstructure:"firstSeed"`
	AccessToken  Token  `mapstructure:"accessToken"`
	RefreshToken Token  `mapstructure:"refreshToken"`
}

type Token struct {
	DefaultTokenFormat string `mapstructure:"defaultFormat"`
	Lifetime           int    `mapstructure:"lifetime"`
}

func LoadConfigFile(profile string) (*Config, error) {
	var err error
	var serverConfig Config
	configReader.AddConfigPath(".")
	configReader.SetConfigFile("config.yaml")
	if err := configReader.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// config file not found; ignore error if desired
			log.Println("file not found", err)
			return nil, err
		} else {
			// config file was found but another error was produced
			log.Println("file open error", err)
			return nil, err
		}
	}
	sub := configReader.Sub(profile)
	if err = sub.Unmarshal(&serverConfig); err != nil {
		log.Println("unmarshal error: ", err)
		return nil, err
	}
	return &serverConfig, nil
}
