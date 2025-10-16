package task

import "github.com/furqanrupom/sqlx-macro/domain"

type service struct {
	taskRepo TaskRepo
}

func NewTaskRepo(taskRepo TaskRepo) Service {
	return &service{
		taskRepo: taskRepo,
	}
}

func (s *service) List() ([]*domain.Task, error) {
	return s.taskRepo.List()
}
func (s *service) Create(task domain.Task) (*domain.Task, error) {
	return s.taskRepo.Create(task)
}
func (s *service) Get(ID int) (*domain.Task, error) {
	return s.taskRepo.Get(ID)
}
func (s *service) Update(ID int) (*domain.Task, error) {
	return s.taskRepo.Update(ID)
}
func (s *service) Delete(ID int) (*domain.Task, error) {
	return s.taskRepo.Delete(ID)
}
