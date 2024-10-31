package users

import (
	"errors"
	"go-skeleton/internal/application/domain/users"
	"go-skeleton/internal/application/services"
)

type Data struct {
	Name  string `db:"Name"`
	Email string `db:"Email"`
}

type Request struct {
	Id     string
	Data   *Data
	Domain *users.Users

	validator services.Validator
}

func NewRequest(id string, data *Data, validator services.Validator) Request {
	domain := &users.Users{}
	return Request{
		Data:      data,
		Domain:    domain,
		Id:        id,
		validator: validator,
	}
}

func (r *Request) Validate() error {
	if r.Id == "" {
		return errors.New("invalid id")
	}

	errs := r.validator.ValidateStruct(r.Data)
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
