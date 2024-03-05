package model

type IntrospectionRequest struct {
	Token         string
	TokenTypeHint string // AccessToken || Refresh Token
}
