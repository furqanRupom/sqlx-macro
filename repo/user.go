package repo

import (
	"errors"

	"github.com/furqanrupom/sqlx-macro/domain"
	"github.com/furqanrupom/sqlx-macro/user"
	"github.com/jmoiron/sqlx"
	"github.com/furqanrupom/sqlx-macro/pkg"
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
	var exists bool
	err := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`, user.Email).Scan(&exists)
	if err != nil {
		return nil,err
	}
	if  exists {
		return nil, errors.New("user already Exits")
	}

	hashPassword,err := pkg.HashPassword(user.Password,pkg.Params)
	if err != nil {
		return nil, errors.New("password hashing failed")
	}
	user.Password = hashPassword
	err = r.db.QueryRow(
    `INSERT INTO users (name, email, password)
     VALUES ($1, $2, $3)
     RETURNING id`,
    user.Name, user.Email, user.Password,
   ).Scan(&user.ID)

   if err !=  nil {
	return nil, errors.New("user registaiton failed")
   }
  return &user,nil
}

func (r *userRepo) Login(email string, password string)(*domain.User,error) {
	var oldUser domain.User
	  err := r.db.QueryRow(
        `SELECT id, name, email, password FROM users WHERE email=$1`,
        email,
    ).Scan(&oldUser.ID, &oldUser.Name, &oldUser.Email, &oldUser.Password)
	if err != nil {
		return nil,err
	}


	compareHash,err := pkg.ComparePassword(password,oldUser.Password)
	if err != nil {
		return nil, err
	}
	if !compareHash {
		return nil, errors.New("wrong password")
	}
	return &oldUser,nil
	
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