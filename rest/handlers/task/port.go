package task

import "github.com/furqanrupom/sqlx-macro/domain"

type Service interface {
	List() ([]*domain.Task, error)
	Create(Task domain.Task) (*domain.Task, error)
	Get(ID int) (*domain.Task, error)
	Update(ID int) (*domain.Task, error)
	Delete(ID int) (*domain.Task, error)
}
