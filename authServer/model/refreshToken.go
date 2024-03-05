package model

import "time"

type RefreshToken struct {
	RefreshToken string
	ExpiresAt    time.Time
	Subject      string
	ClientdId    string
	Issuer       string
	Scope        []string
	CreatedAt    time.Time
	NotBefore    time.Time
	Revoked      bool
}
