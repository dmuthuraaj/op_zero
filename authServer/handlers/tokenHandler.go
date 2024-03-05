package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/op_zero/authserver/model"
)

func (ah *AuthHandler) TokenHandler(c *gin.Context) {
	var err error
	// TODO: Query parameter helper(GET) function & Validation for all Queries
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized client"})
		return
	}
	tokenRequest := &model.OauthTokenRequest{}
	log.Println("username: ", username)
	tokenRequest.ClientId = username
	tokenRequest.ClientSecret = password
	tokenRequest.GrantType, ok = c.GetQuery("grant_type")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "must add grant_type"})
		return
	}
	scope, ok := c.GetQuery("scope")
	if !ok {
		tokenRequest.Scope = append(tokenRequest.Scope, model.SCOPE_READ)
	}
	tokenRequest.Scope = strings.Split(scope, ":")
	accessToken := &model.OauthAccessTokenResponse{}
	if tokenRequest.GrantType == model.GRANT_TYPE_CLIENT_CREDENTIALS {
		accessToken, err = ah.TokenService.CreateAccessTokenForClientCredentials(c, tokenRequest)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"data": accessToken}})
}
