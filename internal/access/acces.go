package access

import "github.com/moguchev/meloman/internal/models"

const servicePath = "/github.moguchev.meloman.Meloman/"

func AccessibleRoles() map[string][]string {
	return map[string][]string{
		servicePath + "Ping":        nil,
		servicePath + "Auth":        nil,
		servicePath + "CreateUser":  nil,
		servicePath + "GetUserByID": {models.RoleAdmin.String(), models.RoleUser.String()},
		servicePath + "GetUsers":    {models.RoleAdmin.String(), models.RoleUser.String()},
	}
}
