package routes

import (
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg/registry"

	"github.com/labstack/echo/v4"
)

type UsersRoutes struct {
	hand *handlers.UsersHandlers
}

func NewUsersRoutes(reg *registry.Registry) *UsersRoutes {
	hand := handlers.NewUsersHandlers(reg)
	return &UsersRoutes{
		hand: hand,
	}
}

func (hs *UsersRoutes) DeclarePrivateRoutes(server *echo.Group, apiPrefix string) {
	server.PUT(apiPrefix+"/users/:id", hs.hand.HandleEditUsers)
	server.DELETE(apiPrefix+"/users/:id", hs.hand.HandleDeleteUsers)
	server.GET(apiPrefix+"/users", hs.hand.HandleListUsers)
	server.GET(apiPrefix+"/users/:id", hs.hand.HandleGetUsers)
	server.POST(apiPrefix+"/users", hs.hand.HandleCreateUsers)
}

func (hs *UsersRoutes) DeclarePublicRoutes(server *echo.Group, apiPrefix string) {
}
