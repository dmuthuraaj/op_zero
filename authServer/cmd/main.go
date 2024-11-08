package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func main() {
	var err error
	viper.SetConfigFile("config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Println("unable to config file: ", err)
	}
	port := viper.GetString("server.port")
	log.Println("port: ", port)
	handler := http.NewServeMux()
	handler.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":"+port, handler)
}
