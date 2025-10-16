package cmd

import (
	"github.com/furqanrupom/sqlx-macro/config"
	middleware "github.com/furqanrupom/sqlx-macro/rest/middlewares"
)

func Serve() {
	conf := config.GetConfig()
	middleware := middleware.NewMiddleware(conf)
	/*
	Will add other config
	*/

}