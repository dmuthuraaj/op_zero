package main

import (
	"log"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	mongo "gopkg.in/go-oauth2/mongo.v3"
)

func main() {
	manager := manage.NewDefaultManager()

	// client memory store
	storeConfigs := mongo.NewStoreConfig(10, 5)
	mongoConf := mongo.NewConfigNonReplicaSet(
		"mongodb://localhost:27017",
		"oauth2",  // database name
		"root",    // username to authenticate with db
		"example", // password to authenticate with db
		"serviceName",
	)

	// use mongodb token store
	manager.MapTokenStorage(
		mongo.NewTokenStore(mongoConf, storeConfigs), // with timeout
		// mongo.NewTokenStore(mongoConf), // no timeout
	)

	clientStore := mongo.NewClientStore(mongoConf, storeConfigs) // with timeout
	// clientStore := mongo.NewClientStore(mongoConf) // no timeout

	manager.MapClientStorage(clientStore)

	// register a service
	clientStore.Create(&models.Client{
		ID:     "confidential",
		Secret: "demo",
		Domain: "http://localhost",
		UserID: "frontend",
		Public: true,
	})

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}
