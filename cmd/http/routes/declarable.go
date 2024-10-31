package routes

import (
	"go-skeleton/pkg/registry"

	"github.com/labstack/echo/v4"
)

type Declarable interface {
	DeclarePrivateRoutes(server *echo.Group, apiPrefix string)
	DeclarePublicRoutes(server *echo.Group, apiPrefix string)
}

func GetRoutes(reg *registry.Registry) map[string]Declarable {
	health := NewHealthRoute()
	usersListRoutes := NewUsersRoutes(reg)
	loginRoute := NewLoginRoutes(reg)
	enforceRoute := NewEnforceRoutes(reg)
	rolesRoute := NewRolesRoutes(reg)
	//{{codeGen1}}
	routes := map[string]Declarable{
		"health": health,
		"users": usersListRoutes,
		"login": loginRoute,
		"enforce": enforceRoute,
		"roles": rolesRoute,
		//{{codeGen2}}
	}
	return routes
}
