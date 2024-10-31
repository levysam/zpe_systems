package routes

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg/registry"
)

type RolesRoutes struct {
	hand *handlers.RolesHandlers
}

func NewRolesRoutes(reg *registry.Registry) *RolesRoutes {
	hand := handlers.NewRolesHandlers(reg)
	return &RolesRoutes{
		hand: hand,
	}
}

func (hs *RolesRoutes) DeclarePublicRoutes(server *echo.Group, apiPrefix string) {}

func (hs *RolesRoutes) DeclarePrivateRoutes(server *echo.Group, apiPrefix string) {
	server.GET(apiPrefix+"/roles", hs.hand.HandleListRoles)
	server.POST(apiPrefix+"/roles/add_resource_permission", hs.hand.HandleRoles)
	server.POST(apiPrefix+"/roles/setToUser", hs.hand.HandleSetRoleToUser)
	server.DELETE(apiPrefix+"/roles/:role", hs.hand.HandleDeleteRoles)
	server.DELETE(apiPrefix+"/roles/deletePermission/:role/:resource", hs.hand.HandleDeleteRoles)
	server.DELETE(apiPrefix+"/roles/deleteFromUser/:role/:userId", hs.hand.HandleDeleteRoleFromUser)

}
