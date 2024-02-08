package main

import (
	"log"
	"os"

	"github.com/dmuthuraaj/op_zero/usermgt"
)

const (
	PORT = "8083"
)

func main() {
	var err error
	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = PORT
	}
	server := usermgt.NewServer()
	err = server.Run(":" + port)
	if err != nil {
		log.Fatalf("unable to start lisening to server on port: %s", port)
	}
}
