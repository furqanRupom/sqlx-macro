package cmd

import (
	"fmt"

	"github.com/furqanrupom/sqlx-macro/config"
	"github.com/furqanrupom/sqlx-macro/infra/db"
	"github.com/furqanrupom/sqlx-macro/repo"
	"github.com/furqanrupom/sqlx-macro/rest"
	taskHandler "github.com/furqanrupom/sqlx-macro/rest/handlers/task"
		userHandler "github.com/furqanrupom/sqlx-macro/rest/handlers/user"
	middleware "github.com/furqanrupom/sqlx-macro/rest/middlewares"
	"github.com/furqanrupom/sqlx-macro/task"
	"github.com/furqanrupom/sqlx-macro/user"
)

func Serve() {
	conf := config.GetConfig()
	dbCon,err := db.NewConnection(conf.DBConfig)
	if err != nil {
		fmt.Println(err)
	}
	if dbCon == nil {
		fmt.Println(err)
	}
	if dbCon != nil {
		fmt.Println("Database connected successfully!")
	}
	err = db.MigrateDB(dbCon,"./migrations")
	if err != nil {
		fmt.Println(err)
	}

	
	middleware := middleware.NewMiddleware(conf)

	taskRepo := repo.NewTaskRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	taskSvc := task.NewService(taskRepo)
	userSvc := user.NewService(userRepo)

	tskHandler := taskHandler.NewHandler(middleware,taskSvc)
	usrHandler := userHandler.NewHandler(middleware,userSvc)

	server := rest.NewServer(conf,usrHandler,tskHandler)
	server.Start()


	

}