package handlers

import (
	"github.com/labstack/echo/v4"
	_ "go-skeleton/internal/application/services"
	rolesAddPermission "go-skeleton/internal/application/services/roles/ADD_RESOURCE_PERMISSION_TO_ROLE"
	rolesDeleteRole "go-skeleton/internal/application/services/roles/DELETE_ROLE"
	rolesList "go-skeleton/internal/application/services/roles/LIST"
	rolesDeletePermissionFromRole "go-skeleton/internal/application/services/roles/REMOVE_RESOURCE_PERMISSION_FROM_ROLE"
	rolesRemoveFromUser "go-skeleton/internal/application/services/roles/REMOVE_ROLE_FROM_USER"
	rolesSetToUser "go-skeleton/internal/application/services/roles/SET_ROLE_TO_USER"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/roles"
	"net/http"
)

type RolesHandlers struct {
	logger       *logger.Logger
	rolesManager *roles.CasbinRule
}

func NewRolesHandlers(reg *registry.Registry) *RolesHandlers {
	return &RolesHandlers{
		logger:       reg.Inject("logger").(*logger.Logger),
		rolesManager: reg.Inject("roles").(*roles.CasbinRule),
	}
}

// HandleRoles Roles
// @Summary      Roles
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Success      200  {object}  rolesAddPermission.Response
// @Param        Authorization header string true "JWT Auth"
// @Param        request body rolesAddPermission.Data true "body model"
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /roles/add_resource_permission [post]
func (hs *RolesHandlers) HandleRoles(context echo.Context) error {
	s := rolesAddPermission.NewService(hs.logger, hs.rolesManager)

	data := new(rolesAddPermission.Data)
	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	request := rolesAddPermission.NewRequest(data)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(422, err)
	}
	return context.JSON(200, response)
}

// HandleDeleteRoles Roles
// @Summary      Roles
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Success      200  {object}  rolesAddPermission.Response
// @Param        Authorization header string true "JWT Auth"
// @Param        role path string true "Role Name"
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /roles/{role} [DELETE]
func (hs *RolesHandlers) HandleDeleteRoles(context echo.Context) error {
	s := rolesDeleteRole.NewService(hs.logger, hs.rolesManager)

	data := new(rolesDeleteRole.Data)
	data.RoleName = context.Param("role")
	request := rolesDeleteRole.NewRequest(data)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(422, err)
	}
	return context.JSON(200, response)
}

// HandleDeletePermission Roles
// @Summary      Roles
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Success      200  {object}  rolesAddPermission.Response
// @Param        Authorization header string true "JWT Auth"
// @Param        role path string true "Role Name"
// @Param        resource path string true "Resource Name"
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /roles/deletePermission/{role}/{resource} [DELETE]
func (hs *RolesHandlers) HandleDeletePermission(context echo.Context) error {
	s := rolesDeletePermissionFromRole.NewService(hs.logger, hs.rolesManager)

	data := new(rolesDeletePermissionFromRole.Data)
	data.RoleName = context.Param("role")
	data.ResourceName = context.Param("resource")

	request := rolesDeletePermissionFromRole.NewRequest(data)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(422, err)
	}
	return context.JSON(200, response)
}

// HandleSetRoleToUser Roles
// @Summary      Roles
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Success      200  {object}  rolesAddPermission.Response
// @Param        Authorization header string true "JWT Auth"
// @Param        request body rolesSetToUser.Data true "body model"
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /roles/setToUser [POST]
func (hs *RolesHandlers) HandleSetRoleToUser(context echo.Context) error {
	s := rolesSetToUser.NewService(hs.logger, hs.rolesManager)

	data := new(rolesSetToUser.Data)
	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}

	request := rolesSetToUser.NewRequest(data)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(422, err)
	}
	return context.JSON(200, response)
}

// HandleDeleteRoleFromUser Roles
// @Summary      Roles
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Success      200  {object}  rolesAddPermission.Response
// @Param        Authorization header string true "JWT Auth"
// @Param        role path string true "Role Name"
// @Param        userId path string true "User id"
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /roles/deleteFromUser/{role}/{userId} [DELETE]
func (hs *RolesHandlers) HandleDeleteRoleFromUser(context echo.Context) error {
	s := rolesRemoveFromUser.NewService(hs.logger, hs.rolesManager)

	data := new(rolesRemoveFromUser.Data)
	data.RoleName = context.Param("role")
	data.UserId = context.Param("userId")

	request := rolesRemoveFromUser.NewRequest(data)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(422, err)
	}
	return context.JSON(200, response)
}

// HandleListRoles Get Roles
// @Summary      Get Roles
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "JWT Auth"
// @Success      200  {object}  rolesList.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /roles [get]
func (hs *RolesHandlers) HandleListRoles(context echo.Context) error {
	s := rolesList.NewService(hs.logger, hs.rolesManager)
	data := new(rolesList.Data)

	request := rolesList.NewRequest(data)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}
