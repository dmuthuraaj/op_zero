package model

import "time"

const (
	ROLE_SYSADMIN  = "ROLE_SYSADMIN"
	ROLE_ADMIN     = "ROLE_ADMIN"
	ROLE_USER      = "ROLE_USER"
	STATUS_ONLINE  = "online"
	STATUS_OFFLINE = "offline"
)

// Have to Add one sysadmin user.
// User can have different role(admin,user).
type User struct {
	Identifier   string    `json:"identifier" bson:"_id"`
	UserName     string    `json:"userName" bson:"userName" binding:"required"`
	FirstName    string    `json:"firstName" bson:"firstName" binding:"required"`
	LastName     string    `json:"lastName" bson:"lastName"`
	MobileNumber string    `json:"mobileNumber" bson:"mobileNumber" binding:"required"`
	Email        string    `json:"email" bson:"email" binding:"required"`
	Password     string    `json:"password" bson:"password" binding:"required"`
	Tenant       string    `json:"tenant" bson:"tenant" binding:"required"`
	ProfileUrl   string    `json:"profileUrl" bson:"profileUrl"`
	Active       bool      `json:"active" bson:"active"`
	Roles        []string  `json:"roles" bson:"roles" binding:"required"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updatedAt"`
}
