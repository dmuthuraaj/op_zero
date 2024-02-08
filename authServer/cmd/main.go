package main

import (
	"log"
	"os"

	"github.com/op_zero/authserver"
)

const (
	PORT = "8081"
)

func main() {
	var err error
	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = PORT
	}
	server := authserver.NewServer()
	err = server.Run(":" + port)
	if err != nil {
		log.Fatalf("unable to start lisening to server on port: %s", port)
	}
}
