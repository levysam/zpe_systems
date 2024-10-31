package users

import (
	"errors"
	"go-skeleton/internal/application/providers/filters"
	"go-skeleton/internal/application/domain/users"
)

type Request struct {
	Data *Data
	Domain *users.Users
	Filters *filters.Filters
}

type Data struct {
	Page int
	ID string
}

func NewRequest(data *Data, filters *filters.Filters) Request {
	domain := &users.Users{}
	return Request{
		Data: data,
		Filters: filters,
		Domain: domain,
	}
}

func (r *Request) Validate() error {
	if r.Data.Page <= 0 {
		return errors.New("invalid page")
	}

	return nil
}
