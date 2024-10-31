package login

import (
	"github.com/golang-jwt/jwt/v5"
	"go-skeleton/internal/application/domain/users"
	"go-skeleton/internal/application/services"
	"time"
)

type Service struct {
	services.BaseService
	response    *Response
	crypt       services.ICrypt
	repository  users.Repository
	signer      services.Signer
	permissions services.PermissionsManager
}

func NewService(log services.Logger, repository users.Repository, crypt services.ICrypt, signer services.Signer, manager services.PermissionsManager) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		crypt:       crypt,
		repository:  repository,
		signer:      signer,
		permissions: manager,
	}
}

func (s *Service) Execute(request Request) {

	if err := request.Validate(); err != nil {
		s.BadRequest(err.Error())
		return
	}

	s.produceResponseRule(request.Data, request.User)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data, domain *users.Users) {
	user, err := s.repository.Get(*domain, "email", data.Email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			s.NotFound("INVALID_EMAIL")
		} else {
			s.InternalServerError("ERROR ON DB", err)
		}
		return
	}

	isValidPass := s.crypt.CompareHash(user.Password, data.Password)
	if !isValidPass {
		s.CustomError(401, "INVALID_PASSWORD")
		return
	}

	roles, rolesErr := s.permissions.ListUserRoles(user.Id)
	if rolesErr != nil {
		s.InternalServerError("ERROR ON DB", rolesErr)
		return
	}
	claims := jwt.MapClaims{
		"aud": roles,
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	signing, generatedErr := token.SigningString()
	if generatedErr != nil {
		s.CustomError(401, "Error on Sign")
		return
	}
	signed, signErr := s.signer.Sign(signing)

	if signErr != nil {
		s.CustomError(401, "Error on Sign")
		return
	}

	s.response = &Response{
		AccessToken: signed,
	}
}
