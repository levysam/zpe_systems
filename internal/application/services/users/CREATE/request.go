package users

import (
	"go-skeleton/internal/application/services"

	"go-skeleton/internal/application/domain/users"
)

type Data struct {
	Id       string   `db:"ID"`
	Name     string   `db:"Name" validate:"required"`
	Email    string   `db:"Email" validate:"required"`
	Roles    []string `db:"Role"`
	Password string   `db:"Password" validate:"required"`
}

type Request struct {
	Data   *Data
	Domain *users.Users
	Err    error

	validator services.Validator
}

func NewRequest(data *Data, validator services.Validator) Request {
	domain := &users.Users{}
	return Request{
		Data:      data,
		Domain:    domain,
		validator: validator,
	}
}

func (r *Request) Validate() error {
	errs := r.validator.ValidateStruct(r.Data)
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
