package model

import "time"

const (
	STATE_ACTIVE    = "ACTIVE"
	STATE_INACTIVE  = "INACTIVE"
	STATE_SUSPENDED = "SUSPENDED"
)

// Tenant is the Group of users.
// Users will be created according to tenant.
type Tenant struct {
	Identifier  string      `json:"identifier" bson:"_id" binding:"-"`
	DisplayName string      `json:"displayName" bson:"displayName" binding:"required"`
	State       string      `json:"state" bson:"state"`
	TenantName  string      `json:"tenantName" bson:"tenantName" binding:"required"`
	ContactInfo contactInfo `json:"contactInfo" bson:"contactInfo" binding:"required"`
	CreatedAt   time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt" bson:"updatedAt"`
}

type contactInfo struct {
	Name         string `json:"name" bson:"name" binding:"required"`
	Email        string `json:"email" bson:"email" binding:"required,email"`
	MobileNumber string `json:"mobileNumber" bson:"mobileNumber" binding:"required,e164"`
}
