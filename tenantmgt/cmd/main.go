package main

import (
	"log"
	"os"

	"github.com/dmuthuraaj/op_zero/tenantmgt"
)

const (
	PORT = "8082"
)

func main() {
	var err error
	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = PORT
	}
	server := tenantmgt.NewServer()
	err = server.Run(":" + port)
	if err != nil {
		log.Fatalf("unable to start lisening to server on port: %s", port)
	}
}
