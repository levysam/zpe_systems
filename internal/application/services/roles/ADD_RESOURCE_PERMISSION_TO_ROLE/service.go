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
	actionString := ""

	for index, action := range data.Actions {
		if index == 0 {
			actionString = "(" + action + ")"
			continue
		}
		actionString = actionString + "|(" + action + ")"
	}
	if len(data.Actions) == 1 {
		actionString = data.Actions[0]
	}

	changed, addErr := s.manager.AddPermissionToRole("role_"+data.Role, data.Resource, actionString)
	if addErr != nil {
		s.InternalServerError("Error adding permission to role", addErr)
		return
	}
	if !changed {
		s.BadRequest("No changes detected")
	}

	s.response = &Response{
		Result: "The role: " + data.Role + " now has permission to " + data.Resource + " with action/s: " + actionString,
	}
}
