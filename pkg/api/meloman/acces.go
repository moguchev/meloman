package meloman

const servicePath = "/github.moguchev.meloman.Meloman/"

func AccessibleRoles() map[string][]string {

	return map[string][]string{
		servicePath + "Ping":       nil,
		servicePath + "Auth":       nil,
		servicePath + "CreateUser": {"admin", "user"},
		servicePath + "GetUser":    {"admin", "user"},
	}
}
