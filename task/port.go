package task

import (
	"github.com/furqanrupom/sqlx-macro/domain"
	taskHandler "github.com/furqanrupom/sqlx-macro/rest/handlers/task"
)

type Service interface {
	taskHandler.Service
}

type TaskRepo interface {
	List() ([]*domain.Task, error)
	Create(Task domain.Task) (*domain.Task, error)
	Get(ID int) (*domain.Task, error)
	Update(domain.Task) (*domain.Task, error)
	Delete(ID int) (*domain.Task, error)
}
