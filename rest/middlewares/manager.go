package middlewares

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (m *Manager) Use(middleware ...Middleware){
	m.globalMiddlewares = append(m.globalMiddlewares,middleware...)
}
func (m *Manager) With (h http.Handler, middleware ...Middleware) http.Handler {
   for _,middleware := range middleware {
	h = middleware(h)
   }
   return h
}
func (m *Manager) WrapMux (h http.Handler) http.Handler {
   for _,middleware := range m.globalMiddlewares {
	h = middleware(h)
   }
   return h
}