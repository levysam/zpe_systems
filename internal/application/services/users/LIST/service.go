package users

import (
	"go-skeleton/internal/application/domain/users"
	"go-skeleton/internal/application/services"
)

type Service struct {
	services.BaseService
	response     *Response
	repository   users.Repository
	pagProv      users.PaginationProvider
	rolesManager services.PermissionsManager
}

func NewService(log services.Logger, repository users.Repository, pagProv users.PaginationProvider, manager services.PermissionsManager) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		repository:   repository,
		pagProv:      pagProv,
		rolesManager: manager,
	}
}

func (s *Service) Execute(request Request) {
	if err := request.Validate(); err != nil {
		s.BadRequest(err.Error())
		return
	}
	if err := request.SetFiltersRules(); err != nil {
		s.BadRequest(err.Error())
		return
	}
	request.Domain.SetFilters(request.Filters)

	s.produceResponseRule(request.Data.Page, 25, request.Domain)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(page int, limit int, domain *users.Users) {
	err, pagination := s.pagProv.PaginationHandler(*domain, page, limit)
	if err != nil {
		s.CustomError(err.Status, err)
		return
	}

	usersData := []DataResponse{}
	for _, user := range *pagination.Data {
		roles, rolesErr := s.rolesManager.ListUserRoles(user.Id)
		if rolesErr != nil {
			s.InternalServerError("Role Service Error", rolesErr)
			return
		}
		usersData = append(usersData, DataResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Roles: roles,
		})
	}

	s.response = &Response{
		CurrentPage: page,
		TotalPages:  pagination.TotalPages,
		Data:        usersData,
	}
}
