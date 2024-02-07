package model

import "time"

const (
	ROLE_SYSADMIN = "ROLE_SYSADMIN"
	ROLE_ADMIN    = "ROLE_ADMIN"
	ROLE_USER     = "ROLE_USER"
)

// Have to Add one sysadmin user.
// User can have different role(admin,user).
type User struct {
	Identifier    string
	FirstName     string
	LastName      string
	MobileNumber  string
	Email         string
	Password      string
	ProfileUrl    string
	Active        bool
	Tenant        string
	Roles         []string
	CreatedAt     *time.Time
	LasModifiedAt *time.Time
}
