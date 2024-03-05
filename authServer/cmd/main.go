package main

import (
	"log"
	"os"

	"github.com/op_zero/authserver"
	"github.com/op_zero/authserver/config"
)

const (
	PORT = "8090"
)

func main() {
	var err error
	var serverConfig *config.Config
	profile := os.Getenv("PROFILE")
	if profile == "" {
		profile = "dev"
	}
	serverConfig, err = config.LoadConfigFile(profile)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Config: ", serverConfig)
	server, err := authserver.NewServer(*serverConfig)
	if err != nil {
		log.Println("error: ", err)
		return
	}
	err = server.Run(":" + serverConfig.Server.Port)
	if err != nil {
		log.Fatalf("unable to start lisening to server on port: %s", serverConfig.Server.Port)
	}
}
