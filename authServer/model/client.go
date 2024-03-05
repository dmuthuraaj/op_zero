package model

import "time"

// Client For Getting tokens
type Client struct {
	// Id of the Client
	Identifier string `json:"identifier" bson:"_id"`
	// ClientId as well as used as username
	ClientId string `json:"clientId" bson:"clientId"`
	// ClientSecret used as a password
	ClientSecret      string    `json:"clientSecret" bson:"clientSecret"`
	Confidential      bool      `json:"confidential" bson:"confidential"`
	AccessTokenFormat string    `json:"accessTokenFormat" bson:"accessTokenFormat"`
	GrantTypes        []string  `json:"grantTypes" bson:"grantTypes"`
	RedirectUris      []string  `json:"redirectUris" bson:"redirectUris"`
	CorsUris          []string  `json:"corsUris" bson:"corsUris"`
	CreatedAt         time.Time `json:"createdAt" bson:"createdAt"`
	LastModifiedAt    time.Time `json:"lastModifiedAt" bson:"lastModifiedAt"`
}
