package roles

import (
	"go-skeleton/internal/application/services"
)

type Service struct {
	services.BaseService
	response *Response
	manager  services.PermissionsManager
}

func NewService(log services.Logger, manager services.PermissionsManager) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		manager: manager,
	}
}

func (s *Service) Execute(request Request) {

	if err := request.Validate(); err != nil {
		s.BadRequest(err.Error())
		return
	}

	s.produceResponseRule(request.Data)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data) {
	roles, err := s.manager.ListRoles()
	if err != nil {
		s.InternalServerError("error on roles list", err)
		return
	}
	s.response = &Response{
		Roles: roles,
	}

}
