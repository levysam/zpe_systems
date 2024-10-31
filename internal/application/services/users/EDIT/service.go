package users

import (
	"go-skeleton/internal/application/domain/users"
	"go-skeleton/internal/application/services"
)

type Service struct {
	services.BaseService
	response   *Response
	repository users.Repository
}

func NewService(log services.Logger, repository users.Repository) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		repository: repository,
	}
}

func (s *Service) Execute(request Request) {
	if err := request.Validate(); err != nil {
		s.BadRequest(err.Error())
		return
	}

	s.produceResponseRule(request.Id, request.Data, request.Domain)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(id string, data *Data, domain *users.Users) {
	userData, findErr := s.repository.Get(*domain, "ID", id)
	if findErr != nil {
		s.BadRequest("user not found")
		return
	}

	domain.Id = userData.Id
	domain.Name = data.Name
	domain.Email = data.Email
	domain.Password = userData.Password

	affected, err := s.repository.Edit(*domain, "Id", id)
	if err != nil {
		s.InternalServerError("error on edit", err)
		return
	}

	if affected < 1 {
		s.UnprocessableEntity("same data or invalid id")
		return
	}

	s.response = &Response{
		Data: data,
	}
}
