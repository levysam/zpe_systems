package routes

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg/registry"
)

type LoginRoutes struct {
	hand *handlers.LoginHandlers
}

func NewLoginRoutes(reg *registry.Registry) *LoginRoutes {
	hand := handlers.NewLoginHandlers(reg)
	return &LoginRoutes{
		hand: hand,
	}
}

func (hs *LoginRoutes) DeclarePublicRoutes(server *echo.Group, apiPrefix string) {
	server.POST(apiPrefix+"/login", hs.hand.HandleLogin)
}

func (hs *LoginRoutes) DeclarePrivateRoutes(server *echo.Group, apiPrefix string) {
}
