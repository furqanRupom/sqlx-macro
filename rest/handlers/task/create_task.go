package task

import (
	"net/http"
)

type CreateTask struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Status      string `json:"status" db:"status"`
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {}
