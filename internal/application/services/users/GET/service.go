package users

import (
	"fmt"
	"go-skeleton/internal/application/domain/users"
	"go-skeleton/internal/application/services"
)

type Service struct {
	services.BaseService
	response   *Response
	repository users.Repository
	manager    services.PermissionsManager
}

func NewService(log services.Logger, repository users.Repository, manager services.PermissionsManager) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		repository: repository,
		manager:    manager,
	}
}

func (s *Service) Execute(request Request) {
	if err := request.Validate(); err != nil {
		s.BadRequest(err.Error())
		return
	}
	s.produceResponseRule(request.Data, request.Domain)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data, domain *users.Users) {
	dom, err := s.repository.Get(*domain, "Id", data.Id)

	if err != nil {
		fmt.Println(err)
		if err.Error() == "sql: no rows in result set" {
			s.NotFound("data not found")
			return
		}
		s.InternalServerError("error on get", err)
		return
	}

	roles, err := s.manager.ListUserRoles(dom.Id)
	if err != nil {
		s.InternalServerError("error on list", err)
		return
	}
	s.response = &Response{
		Data: UserResponse{
			Id:    dom.Id,
			Name:  dom.Name,
			Email: dom.Email,
			Roles: roles,
		},
	}
}
