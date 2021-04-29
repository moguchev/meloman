package models

import (
	"strings"

	"github.com/moguchev/meloman/pkg/api/meloman"
)

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

func (ur UserRole) String() string {
	return string(ur)
}

func (ur UserRole) Proto() meloman.Role {
	switch ur {
	case RoleUser:
		return meloman.Role_USER
	case RoleAdmin:
		return meloman.Role_ADMIN
	default:
		return meloman.Role_USER
	}
}

func ParseRole(s string) UserRole {
	switch strings.ToLower(s) {
	case "user":
		return RoleUser
	case "admin":
		return RoleAdmin
	default:
		return RoleUser
	}
}
