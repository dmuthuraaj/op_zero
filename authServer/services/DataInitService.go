package services

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/op_zero/authserver/model"
	"github.com/op_zero/authserver/utils"
)

func (ts *TokenService) DataInitialize() error {
	corsUris := []string{
		"http://localhost:4000/index.html",
	}
	redirectUris := []string{
		"http://localhost:4000/index.html",
	}
	var client1 model.Client
	client1.Identifier = uuid.New().String()
	client1.ClientId = "confidential-jwt"
	client1.ClientSecret, _ = utils.PasswordEncoder("demo")
	client1.AccessTokenFormat = "JWT"
	client1.Confidential = true
	client1.CorsUris = corsUris
	client1.RedirectUris = redirectUris
	client1.GrantTypes = append(client1.GrantTypes, model.GRANT_TYPE_AUTHORIZATION_CODE, model.GRANT_TYPE_CLIENT_CREDENTIALS, model.GRANT_TYPE_PASSWORD, model.GRANT_TYPE_REFRESH_TOKEN)
	client1.CreatedAt = time.Now()
	client1.LastModifiedAt = time.Now()
	log.Println("adding Client")
	err := ts.TokenDatastore.AddClient(context.Background(), client1)
	if err != nil {
		log.Println("adding Client failed with err: ", err)
		return err
	}
	return nil
}
