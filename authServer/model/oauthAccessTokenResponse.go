package model

import (
	"time"
)

type OauthAccessTokenResponse struct {
	Identity     string
	AccessToken  string
	RefreshToken string
	TokenType    string
	// Scope        string
	ExpiresIn int
	CreatedAt time.Time
}

// func NewAccessToken(scope string, tokenType string, expiresIn int) *OauthAccessTokenResponse {
// 	accessToken := &OauthAccessTokenResponse{
// 		AccessToken:  uuid.New().String(),
// 		RefreshToken: uuid.New().String(),
// 		ExpiresIn:    expiresIn,
// 		TokenType:    tokenType,
// 		Scope:        scope,
// 	}
// 	return accessToken
// }
