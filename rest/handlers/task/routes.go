package task

import (
	"net/http"

	middleware "github.com/furqanrupom/sqlx-macro/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /tasks",
		manager.With(
			http.HandlerFunc(h.CreateTask),
		),
	)
	mux.Handle("GET /tasks",
		manager.With(
			http.HandlerFunc(h.TaskList),
		),
	)
}
