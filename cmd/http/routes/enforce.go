package routes

import (
	"github.com/labstack/echo/v4"
	"go-skeleton/cmd/http/handlers"
	"go-skeleton/pkg/registry"
)

type EnforceRoutes struct {
	hand *handlers.EnforceHandlers
}

func NewEnforceRoutes(reg *registry.Registry) *EnforceRoutes {
	hand := handlers.NewEnforceHandlers(reg)
	return &EnforceRoutes{
		hand: hand,
	}
}

func (hs *EnforceRoutes) DeclarePublicRoutes(server *echo.Group, apiPrefix string) {}

func (hs *EnforceRoutes) DeclarePrivateRoutes(server *echo.Group, apiPrefix string) {
	server.POST(apiPrefix+"/enforce", hs.hand.HandleEnforce)
}
