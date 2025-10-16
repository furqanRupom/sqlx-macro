package middlewares

import "github.com/furqanrupom/sqlx-macro/config"

type Middlewares struct {
	config *config.Config
}

func NewMiddleware(config *config.Config) *Middlewares {
	return &Middlewares{
		config : config,
	}
}