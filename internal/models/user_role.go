package models

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

func (ur UserRole) String() string {
	return string(ur)
}
