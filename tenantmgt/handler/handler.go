package handler

import (
	"log"
	"net/http"

	"github.com/dmuthuraaj/usermgt/model"
	"github.com/dmuthuraaj/usermgt/service"
	"github.com/gin-gonic/gin"
)

type TenantHandler struct {
	tenantAdm service.Service
}

func NewTenantHandler(th service.Service) *TenantHandler {
	return &TenantHandler{
		tenantAdm: th,
	}
}

func (th *TenantHandler) CreateTenant(c *gin.Context) {
	var err error
	var tenant model.Tenant
	err = c.ShouldBindJSON(&tenant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = th.tenantAdm.CreateTenant(c, tenant)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": "tenant created successfully"})
}
