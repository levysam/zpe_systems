package usersRepository

import (
	"github.com/jmoiron/sqlx"
	users  "go-skeleton/internal/application/domain/users"
	"go-skeleton/internal/repositories/base_repository"
)

type UsersRepository struct {
	*base_repository.BaseRepo[users.Users]
}

func NewUsersRepository(mysql *sqlx.DB) *UsersRepository {
	return &UsersRepository{
		BaseRepo: base_repository.NewBaseRepository[users.Users](mysql),
	}
}
