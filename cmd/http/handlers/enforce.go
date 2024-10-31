package handlers

import (
	"github.com/labstack/echo/v4"
	_ "go-skeleton/internal/application/services"
	"go-skeleton/internal/application/services/enforce"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/roles"
)

type EnforceHandlers struct {
	logger       *logger.Logger
	rolesManager *roles.CasbinRule
}

func NewEnforceHandlers(reg *registry.Registry) *EnforceHandlers {
	return &EnforceHandlers{
		logger:       reg.Inject("logger").(*logger.Logger),
		rolesManager: reg.Inject("roles").(*roles.CasbinRule),
	}
}

// HandleEnforce Enforce
// @Summary      Enforce
// @Tags         Enforce
// @Accept       json
// @Produce      json
// @Success      200  {object}  enforce.Response
// @Param        request body enforce.Data true "body model"
// @Param        Authorization header string true "JWT Auth"
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /enforce [get]
func (hs *EnforceHandlers) HandleEnforce(context echo.Context) error {
	s := enforce.NewService(hs.logger, hs.rolesManager)

	data := new(enforce.Data)
	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	request := enforce.NewRequest(data)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(422, err)
	}
	return context.JSON(200, response)
}
