package user

import (
	middleware "github.com/furqanrupom/sqlx-macro/rest/middlewares"
)

type Handler struct {
	middleware *middleware.Middlewares
	svc Service
}

func NewHandler(middleware *middleware.Middlewares, svc Service) *Handler {
 return &Handler{
	middleware: middleware,
	svc: svc,
 }
}