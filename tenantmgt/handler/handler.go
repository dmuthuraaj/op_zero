package handler

import (
	"errors"
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

func ErrResponseWithLogger(c *gin.Context, code int, err error) {

	c.JSON(code, gin.H{
		"status": code,
		"error":  err.Error(),
		"path":   c.Request.URL.Path,
	})
}

func (th *TenantHandler) CreateTenant(c *gin.Context) {
	var err error
	var tenant model.Tenant
	c.Writer.Header().Add("X-Request-Id", "1234-5678-9012")
	err = c.ShouldBindJSON(&tenant)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusBadRequest, err)
		return
	}
	err = th.tenantAdm.CreateTenant(c, &tenant)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": "tenant created successfully"})
}

func (th *TenantHandler) GetTenantByName(c *gin.Context) {
	// TODO: Change to Id
	tenantName := c.Param("name")
	if tenantName == "" {
		ErrResponseWithLogger(c, http.StatusBadRequest, errors.New("name required"))
		return
	}
	tenant, err := th.tenantAdm.GetTenantByName(c, tenantName)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &tenant})
}

func (th *TenantHandler) UpdateTenantContactInfo(c *gin.Context) {
	var err error
	var tenantUpdate model.Tenant
	err = c.BindJSON(&tenantUpdate)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusBadRequest, err)
		return
	}
	err = th.tenantAdm.UpdateTenantContactInfo(c, &tenantUpdate)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "tenant contactInfo updated successfully"})
}

func (th *TenantHandler) DeleteTenantByName(c *gin.Context) {
	// TODO: Change to Id
	tenantName := c.Param("name")
	if tenantName == "" {
		ErrResponseWithLogger(c, http.StatusBadRequest, errors.New("name required"))
		return
	}
	err := th.tenantAdm.DeleteTenantByName(c, tenantName)
	if err != nil {
		ErrResponseWithLogger(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"data": "tenant deleted successfully"})
}
