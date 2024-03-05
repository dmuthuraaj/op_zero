package handler

import "github.com/op_zero/authserver/services"

type AuthHandler struct {
	TokenService services.Service
}

func NewAuthHandler(tokenService services.Service) *AuthHandler {
	return &AuthHandler{TokenService: tokenService}
}
