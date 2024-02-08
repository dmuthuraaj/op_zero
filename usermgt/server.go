package usermgt

import (
	"github.com/dmuthuraaj/op_zero/usermgt/datastore"
	"github.com/dmuthuraaj/op_zero/usermgt/handler"
	"github.com/dmuthuraaj/op_zero/usermgt/service"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

const (
	uriUser = "/users"
)

func NewServer() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.SetLogger())

	datastore := datastore.NewUserDatastore()
	service := service.NewUserService(datastore)
	handler := handler.NewUserHandler(service)

	userRouterGroup := r.Group("api/v1")
	// Create user
	userRouterGroup.POST(uriUser, handler.CreateUser)
	// Get All Users
	userRouterGroup.GET(uriUser, handler.GetAllUsers)
	// Get user By Name
	userRouterGroup.GET(uriUser+"/:name", handler.GetUserByName)
	// Delete user By Name
	userRouterGroup.DELETE(uriUser+"/:name", handler.DeleteUserByName)
	// Update user
	userRouterGroup.PUT(uriUser, handler.UpdateUser)
	return r
}
