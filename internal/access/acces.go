package access

import "github.com/moguchev/meloman/internal/models"

const servicePath = "/github.moguchev.meloman.Meloman/"

func AccessibleRoles() map[string][]string {
	return map[string][]string{
		// public space
		servicePath + "Ping":       nil,
		servicePath + "Auth":       nil,
		servicePath + "CreateUser": nil,
		// admin space
		servicePath + "UpdateUserRole": {models.RoleAdmin.String()},
		servicePath + "CreateArtist":   {models.RoleAdmin.String()},
		servicePath + "CreateAlbum":    {models.RoleAdmin.String()},
		servicePath + "CreateTrack":    {models.RoleAdmin.String()},
		// user space
		servicePath + "GetUsers":        {models.RoleAdmin.String(), models.RoleUser.String()},
		servicePath + "GetUserByID":     {models.RoleAdmin.String(), models.RoleUser.String()},
		servicePath + "GetArtists":      {models.RoleAdmin.String(), models.RoleUser.String()},
		servicePath + "GetArtistByID":   {models.RoleAdmin.String(), models.RoleUser.String()},
		servicePath + "GetArtistAlbums": {models.RoleAdmin.String(), models.RoleUser.String()},
		servicePath + "GetFormats":      {models.RoleAdmin.String(), models.RoleUser.String()},
		servicePath + "GetLabels":       {models.RoleAdmin.String(), models.RoleUser.String()},
	}
}
