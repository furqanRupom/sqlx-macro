package repo

import (
	"github.com/furqanrupom/sqlx-macro/domain"
	"github.com/furqanrupom/sqlx-macro/task"
	"github.com/jmoiron/sqlx"
)

type TaskRepo struct {
	task.Service
}

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{
		db: db,
	}
}

func (r *taskRepo) Create(task domain.Task) (*domain.Task, error) {
	return &task, nil
}
func (r *taskRepo) List() ([]*domain.Task, error) {
	return nil, nil
}
func (r *taskRepo) Get(ID int) (*domain.Task, error) {
	return nil, nil
}

func (r *taskRepo) Update(task domain.Task) (*domain.Task, error) {
	return nil, nil
}
func (r *taskRepo) Delete(ID int) (*domain.Task, error) {
	return nil, nil
}
