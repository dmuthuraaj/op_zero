package usermgt

import (
	"github.com/dmuthuraaj/usermgt/datastore"
	"github.com/dmuthuraaj/usermgt/handler"
	"github.com/dmuthuraaj/usermgt/service"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

const (
	uriTenant            = "/tenant"
	uriTenantContactInfo = "/tenant/contact"
)

func NewServer() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.SetLogger())

	datastore := datastore.NewTenantDatastore()
	service := service.NewTenantService(datastore)
	handler := handler.NewTenantHandler(service)

	tenantRouterGroup := r.Group("api/v1")
	// Create Tenant
	tenantRouterGroup.POST(uriTenant, handler.CreateTenant)
	// Get Tenant By Name
	tenantRouterGroup.GET(uriTenant+"/:name", handler.GetTenantByName)
	// Delete Tenant By Name
	tenantRouterGroup.DELETE(uriTenant+"/:name", handler.DeleteTenantByName)
	// Update Tenant ContactInfo
	tenantRouterGroup.PUT(uriTenantContactInfo, handler.UpdateTenantContactInfo)
	return r
}
