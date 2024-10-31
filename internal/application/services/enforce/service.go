package enforce

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
	permission, aclErr := s.manager.CheckPermission(data.UserId, data.Resource, data.Method)
	if aclErr != nil {
		s.InternalServerError("Error on db", aclErr)
		return
	}

	if !permission {
		s.response = &Response{
			Allowed: false,
		}
		return
	}
	s.response = &Response{
		Allowed: true,
	}
}
