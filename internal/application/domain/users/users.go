package users

import (
	"go-skeleton/internal/application/providers/filters"
	"go-skeleton/internal/application/providers/pagination"
	"go-skeleton/internal/repositories/base_repository"
)

type Users struct {
	Id       string `db:"ID"`
	Name     string `db:"Name"`
	Email    string `db:"Email"`
	Password string `db:"Password"`
	client   string
	filters  *filters.Filters
}

func (d *Users) SetClient(client string) {
	d.client = client
}

func (d *Users) SetFilters(filters *filters.Filters) {
	d.filters = filters
}

func (d Users) GetFilters() filters.Filters {
	if d.filters != nil {
		return *d.filters
	}
	return filters.Filters{}
}

func (d Users) Schema() string {
	return "users"
}

type Repository interface {
	base_repository.BaseRepository[Users]
}

type PaginationProvider interface {
	pagination.IPaginationProvider[Users]
}
