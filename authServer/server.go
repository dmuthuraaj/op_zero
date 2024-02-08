package authserver

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

const (
	uriAuth     = "/auth"
	uriAutorize = "/authorize"
)

func NewServer() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.SetLogger())

	// datastore := datastore.NewAuthStore()
	// service := service.NewAuthService(datastore)
	// handler := handler.NewAuthHandler(service)

	// authRouterGroup := r.Group("auth")
	// authRouterGroup.POST("", handler.Login)
	return r
}
