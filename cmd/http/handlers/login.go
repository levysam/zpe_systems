package handlers

import (
	"github.com/labstack/echo/v4"
	_ "go-skeleton/internal/application/services"
	"go-skeleton/internal/application/services/login"
	usersRepository "go-skeleton/internal/repositories/users"
	"go-skeleton/pkg/crypt"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/roles"
	"go-skeleton/pkg/signerVerifier"
)

type LoginHandlers struct {
	UsersRepository *usersRepository.UsersRepository

	logger       *logger.Logger
	crypto       *crypt.Crypt
	signer       *signerVerifier.Signer
	rolesManager *roles.CasbinRule
}

func NewLoginHandlers(reg *registry.Registry) *LoginHandlers {
	return &LoginHandlers{
		logger:          reg.Inject("logger").(*logger.Logger),
		crypto:          reg.Inject("crypt").(*crypt.Crypt),
		UsersRepository: reg.Inject("usersRepository").(*usersRepository.UsersRepository),
		signer:          reg.Inject("signerVerifier").(*signerVerifier.Signer),
		rolesManager:    reg.Inject("roles").(*roles.CasbinRule),
	}
}

// HandleLogin Login
// @Summary      Login
// @Tags         Login
// @Accept       json
// @Produce      json
// @Success      200  {object}  login.Response
// @Param        request body login.Data true "body model"
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /login [post]
func (hs *LoginHandlers) HandleLogin(context echo.Context) error {
	s := login.NewService(hs.logger, hs.UsersRepository, hs.crypto, hs.signer, hs.rolesManager)

	data := new(login.Data)
	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	request := login.NewRequest(data)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(422, err)
	}
	return context.JSON(200, response)
}
