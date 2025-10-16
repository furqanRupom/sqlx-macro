package user

import "github.com/furqanrupom/sqlx-macro/domain"

type service struct {
	userRepo UserRepo
}

func NewUserRepo(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Register(user domain.User) (*domain.User, error) {
	return s.userRepo.Register(user)
}
func (s *service) Login(email string, password string) (string, error) {
	return s.userRepo.Login(email, password)
}
func (s *service) Get(ID int) (*domain.User, error) {
	return s.userRepo.Get(ID)
}
func (s *service) Delete(ID int) (*domain.User, error) {
	return s.userRepo.Delete(ID)
}
func (s *service) Update(ID int) (*domain.User, error) {
	return s.userRepo.Update(ID)
}
