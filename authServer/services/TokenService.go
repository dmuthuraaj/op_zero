package services

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/op_zero/authserver/config"
	"github.com/op_zero/authserver/datastore"
	"github.com/op_zero/authserver/model"
	"github.com/op_zero/authserver/utils"
)

type TokenService struct {
	TokenDatastore datastore.Datastore
}

func NewTokenService(tokenDatastore datastore.Datastore, config config.Config) *TokenService {
	return &TokenService{
		TokenDatastore: tokenDatastore,
	}
}

func (ts *TokenService) CreateAccessTokenForClientCredentials(ctx context.Context, tokenReq *model.OauthTokenRequest) (*model.OauthAccessTokenResponse, error) {
	var err error
	var tokenReponse *model.OauthAccessTokenResponse
	client, err := ts.TokenDatastore.GetClientByClientId(ctx, tokenReq.ClientId)
	if err != nil {
		log.Println("getting client failed with error: ", err)
		return nil, err
	}
	err = utils.ComparePassword(tokenReq.ClientSecret, client.ClientSecret)
	if err != nil {
		log.Println("client password is wrong: ", err)
		return nil, err
	}
	accessToken, err := createAnonymousAccessToken(*tokenReq)
	if err != nil {
		log.Println("creating AccessToken Failed with error: ", err)
		return nil, err
	}
	refreshToken, err := createAnonymousRefreshAccessToken(*tokenReq)
	if err != nil {
		log.Println("creating RefreshToken Failed with error: ", err)
		return nil, err
	}
	err = ts.TokenDatastore.SaveRefreshToken(ctx, refreshToken)
	if err != nil {
		log.Println("saving refreshToken Failed with error: ", err)
		return nil, err
	}
	tokenReponse.Identity = uuid.NewString()
	tokenReponse.AccessToken = accessToken
	tokenReponse.TokenType = model.TOKEN_TYPE_BEARER
	tokenReponse.ExpiresIn = int(time.Now().Add(24 * time.Hour).Unix())
	tokenReponse.CreatedAt = time.Now().UTC()
	// tokenReponse.Scope = tokenReq.Scope
	tokenReponse.RefreshToken = refreshToken.RefreshToken
	err = ts.TokenDatastore.SaveAccessToken(ctx, tokenReponse)
	if err != nil {
		log.Println("saving AccessToken Failed with error: ", err)
		return nil, err
	}
	return tokenReponse, nil
}

func createAnonymousAccessToken(tokenReq model.OauthTokenRequest) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "authserver",
		Subject:   tokenReq.ClientId,
		Audience:  []string{tokenReq.ClientId},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		ID:        uuid.NewString(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	privKey, err := utils.ReadPrivateKey()
	if err != nil {
		return "", nil
	}
	accessToken, err := token.SignedString(privKey)
	if err != nil {
		log.Println("accessToken error: ", accessToken)
		return "", err
	}
	log.Println("accessToken: ", accessToken)
	return accessToken, nil
}

func createPersonalizedJWTToken() (string, error) {
	return "", nil
}

func parseAndValidateJWTToken(token string) error {
	return nil
}

func createAnonymousRefreshAccessToken(tokenReq model.OauthTokenRequest) (*model.RefreshToken, error) {
	var refreshToken model.RefreshToken
	refreshToken.RefreshToken = utils.RandString(48)
	refreshToken.Subject = "anonymous"
	refreshToken.CreatedAt = time.Now().UTC()
	refreshToken.ExpiresAt = time.Now().Add(24 * time.Hour).UTC()
	refreshToken.ClientdId = tokenReq.ClientId
	refreshToken.Scope = tokenReq.Scope
	refreshToken.Issuer = "authserver"
	refreshToken.NotBefore = time.Now().UTC()
	return &refreshToken, nil
}
