package handlers

import (
	"go-skeleton/internal/application/domain/users"
	"go-skeleton/internal/application/providers/filters"
	"go-skeleton/internal/application/providers/pagination"
	_ "go-skeleton/internal/application/services"
	usersCreate "go-skeleton/internal/application/services/users/CREATE"
	usersDelete "go-skeleton/internal/application/services/users/DELETE"
	usersEdit "go-skeleton/internal/application/services/users/EDIT"
	usersGet "go-skeleton/internal/application/services/users/GET"
	usersList "go-skeleton/internal/application/services/users/LIST"
	usersRepository "go-skeleton/internal/repositories/users"
	"go-skeleton/pkg/crypt"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/roles"
	"go-skeleton/pkg/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersHandlers struct {
	UsersRepository *usersRepository.UsersRepository

	logger       *logger.Logger
	idCreator    *idCreator.IdCreator
	validator    *validator.Validator
	crypto       *crypt.Crypt
	rolesManager *roles.CasbinRule
}

func NewUsersHandlers(reg *registry.Registry) *UsersHandlers {
	return &UsersHandlers{
		UsersRepository: reg.Inject("usersRepository").(*usersRepository.UsersRepository),
		logger:          reg.Inject("logger").(*logger.Logger),
		idCreator:       reg.Inject("idCreator").(*idCreator.IdCreator),
		validator:       reg.Inject("validator").(*validator.Validator),
		crypto:          reg.Inject("crypt").(*crypt.Crypt),
		rolesManager:    reg.Inject("roles").(*roles.CasbinRule),
	}
}

// HandleGetUsers Get Users
// @Summary      Get a Users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        users_id path string true "Users ID"
// @Param        Authorization header string true "JWT Auth"
// @Success      200  {object}  usersGet.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /users/{users_id} [get]
func (hs *UsersHandlers) HandleGetUsers(context echo.Context) error {
	s := usersGet.NewService(hs.logger, hs.UsersRepository, hs.rolesManager)
	data := new(usersGet.Data)

	data.Id = context.Param("id")

	request := usersGet.NewRequest(data)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

// HandleCreateUsers Create Users
// @Summary      Create Users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body usersCreate.Data true "body model"
// @Param        Authorization header string true "JWT Auth"
// @Success      200  {object}  usersCreate.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /users [post]
func (hs *UsersHandlers) HandleCreateUsers(context echo.Context) error {
	s := usersCreate.NewService(hs.logger, hs.UsersRepository, hs.idCreator, hs.crypto, hs.rolesManager)
	data := new(usersCreate.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}
	request := usersCreate.NewRequest(data, hs.validator)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusCreated, response)
}

// HandleEditUsers Edit Users
// @Summary      Edit Users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        users_id path string true "Users ID"
// @Param        request body usersEdit.Data true "body model"
// @Param        Authorization header string true "JWT Auth"
// @Success      200  {object}  usersEdit.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /users/{users_id} [put]
func (hs *UsersHandlers) HandleEditUsers(context echo.Context) error {
	s := usersEdit.NewService(hs.logger, hs.UsersRepository)
	data := new(usersEdit.Data)

	id := context.Param("id")
	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}

	request := usersEdit.NewRequest(id, data, hs.validator)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

// HandleListUsers List Users
// @Summary      List Users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        page  query   int  true  "valid int"
// @Param        Authorization header string true "JWT Auth"
// @Success      200  {object}  usersList.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /users [get]
func (hs *UsersHandlers) HandleListUsers(context echo.Context) error {
	s := usersList.NewService(
		hs.logger,
		hs.UsersRepository,
		pagination.NewPaginationProvider[users.Users](hs.UsersRepository),
		hs.rolesManager,
	)

	data := new(usersList.Data)
	bindErr := echo.QueryParamsBinder(context).
		Int("page", &data.Page).
		String("id", &data.ID).
		BindErrors()

	if bindErr != nil {
		return context.JSON(http.StatusBadRequest, bindErr)
	}

	f := filters.NewFilters()

	request := usersList.NewRequest(data, f)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

// HandleDeleteUsers Delete Users
// @Summary      Delete Users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        users_id path string true "Users ID"
// @Param        Authorization header string true "JWT Auth"
// @Success      200  {object}  usersDelete.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /users/{user_id} [DELETE]
func (hs *UsersHandlers) HandleDeleteUsers(context echo.Context) error {
	s := usersDelete.NewService(hs.logger, hs.UsersRepository)
	data := new(usersDelete.Data)

	data.Id = context.Param("id")

	request := usersDelete.NewRequest(data)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}
