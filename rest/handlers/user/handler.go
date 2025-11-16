package user

import (
	"github.com/furqanrupom/sqlx-macro/config"
	middleware "github.com/furqanrupom/sqlx-macro/rest/middlewares"
)

type Handler struct {
	middleware *middleware.Middlewares
	config *config.Config
	svc Service
}

func NewHandler(middleware *middleware.Middlewares, svc Service, config *config.Config) *Handler {
 return &Handler{
	middleware: middleware,
	svc: svc,
	config: config,
 }
}