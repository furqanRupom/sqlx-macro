package user

import (
	"github.com/furqanrupom/sqlx-macro/domain"
	userHandler "github.com/furqanrupom/sqlx-macro/rest/handlers/user"
)

type UserRepo interface {
	userHandler.Service
}

type Service interface {
	Register(u domain.User) (*domain.User, error)
	Login(email string, password string) (string, error)
	Get(ID int) (*domain.User, error)
	Update(domain.User) (*domain.User, error)
	Delete(ID int) (*domain.User, error)
}
