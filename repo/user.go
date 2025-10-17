package repo

import (
	"github.com/furqanrupom/sqlx-macro/domain"
	"github.com/furqanrupom/sqlx-macro/user"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	user.Service
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Register(user domain.User)(*domain.User,error) {
	return nil, nil
}

func (r *userRepo) Login(email string, password string)(string,error) {
	return "", nil
}
func (r *userRepo) Get(ID int)(*domain.User,error) {
	return nil, nil
}
func (r *userRepo) Update(user domain.User)(*domain.User,error) {
	return nil, nil
}
func (r *userRepo) Delete(ID int)(*domain.User,error) {
	return nil, nil
}