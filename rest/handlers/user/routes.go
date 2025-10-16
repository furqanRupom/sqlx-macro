package user

import (
	middleware "github.com/furqanrupom/sqlx-macro/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /register",
		manager.With(
			http.HandlerFunc(h.RegisterUser),
		),
	)
	mux.Handle("POST /login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)

}
