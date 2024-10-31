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
	s.produceResponseRule(request.Data, request.Domain)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data, domain *users.Users) {
	err := s.repository.Delete(*domain, "Id", data.Id)
	if err != nil {
		s.InternalServerError("error on delete", err)
		return
	}

	s.response = &Response{
		Message: "OK",
	}
}
