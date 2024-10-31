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
	changed, err := s.manager.DeletePermissionFromRole("role_"+data.RoleName, data.ResourceName)
	if err != nil {
		s.InternalServerError("Error on delete permission from role", err)
		return
	}
	if !changed {
		s.BadRequest("No permission changed on role")
		return
	}
	s.response = &Response{
		Data: true,
	}
}
