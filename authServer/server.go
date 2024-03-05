package authserver

import (
	"log"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/op_zero/authserver/config"
	"github.com/op_zero/authserver/datastore"
	handler "github.com/op_zero/authserver/handlers"
	"github.com/op_zero/authserver/services"
)

const (
	uriToken    = "/token"
	uriRevoke   = "/revoke"
	uriDevice   = "/device/code"
	uriUserInfo = "/userinfo"
	uriAutorize = "/authorize"
)

func NewServer(config config.Config) (*gin.Engine, error) {
	var err error
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.SetLogger())
	tokenDatastore, err := datastore.NewTokenDatastore(config)
	if err != nil {
		return nil, err
	}
	authService := services.NewTokenService(tokenDatastore, config)
	handler := handler.NewAuthHandler(authService)
	if config.Server.FirstSeed {
		err = authService.DataInitialize()
		if err != nil {
			log.Println("data initialization failed with error: ", err)
			return nil, err
		}
	}
	authRouterGroup := r.Group("auth")
	authRouterGroup.POST(uriToken, handler.TokenHandler)
	return r, nil
}
