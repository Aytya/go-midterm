package domain

type Role struct {
	Name        string
	Permissions []string
}

var Roles = map[string]Role{
	"admin": {
		Name:        "admin",
		Permissions: []string{"create", "read", "update", "delete"},
	},
	"user": {
		Name:        "user",
		Permissions: []string{"read"},
	},
}
