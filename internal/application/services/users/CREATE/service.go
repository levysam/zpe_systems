package users

import (
	"go-skeleton/internal/application/domain/users"
	"go-skeleton/internal/application/services"
)

type Service struct {
	services.BaseService
	response    *Response
	repository  users.Repository
	idCreator   services.IdCreator
	crypt       services.ICrypt
	permissions services.PermissionsManager
}

func NewService(log services.Logger, repository users.Repository, idCreator services.IdCreator, crypt services.ICrypt, manager services.PermissionsManager) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		repository:  repository,
		idCreator:   idCreator,
		crypt:       crypt,
		permissions: manager,
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
	domain.Id = s.idCreator.Create()
	domain.Name = data.Name
	domain.Email = data.Email

	encryptedPass, passErr := s.crypt.GenerateHash(data.Password)
	if passErr != nil {
		s.InternalServerError("error on password encryption", passErr)
		return
	}

	domain.Password = encryptedPass

	tx, txErr := s.repository.InitTX()
	if txErr != nil {
		s.InternalServerError("error on create", txErr)
		return
	}

	err := s.repository.Create(*domain, tx, false)
	if err != nil {
		s.InternalServerError("error on create", err)
		return
	}
	_, setRoleErr := s.permissions.SetRoleToUserBatch(domain.Id, data.Roles)
	if setRoleErr != nil {
		tx.Rollback()
		s.InternalServerError("error on setRole", setRoleErr)
		return
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		s.InternalServerError("error on db", commitErr)
		return
	}

	s.response = &Response{
		Created: true,
	}
}
