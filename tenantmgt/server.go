package usermgt

import (
	"time"

	"github.com/dmuthuraaj/usermgt/datastore"
	"github.com/dmuthuraaj/usermgt/handler"
	"github.com/dmuthuraaj/usermgt/service"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	uriTenant = "/tenant"
)

func NewServer() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r := gin.New()
	// Middleware
	// r.Use(middleware.NewLogger())
	logger, _ := zap.NewProduction()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger, true))

	datastore := datastore.NewTenantDatastore()
	service := service.NewTenantService(datastore)
	handler := handler.NewTenantHandler(service)
	tenantRouterGroup := r.Group("api/v1")
	tenantRouterGroup.POST(uriTenant, handler.CreateTenant)
	return r
}
