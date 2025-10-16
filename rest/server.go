package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/furqanrupom/sqlx-macro/config"
	"github.com/furqanrupom/sqlx-macro/rest/handlers/task"
	"github.com/furqanrupom/sqlx-macro/rest/handlers/user"
	middleware "github.com/furqanrupom/sqlx-macro/rest/middlewares"
)



type Server struct {
	conf        *config.Config
	userHandler *user.Handler
	taskHandler *task.Handler
}

// new Server

func NewServer(
	conf *config.Config,
	userHandler *user.Handler,
	taskHandler *task.Handler,

) *Server {
	return &Server{
		conf:        conf,
		userHandler: userHandler,
		taskHandler: taskHandler,
	}
}
// start server

func (server *Server) Start(){
	// routes and middleware config
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.PreFlight,middleware.Cors,middleware.Logger)
	wrappedMux := manager.WrapMux(mux)

	// handlers
    server.userHandler.RegisterRoutes(mux, manager)
	server.taskHandler.RegisterRoutes(mux,manager)
	addr := "127.0.0.1:" + strconv.Itoa(server.conf.HttpPort)
	fmt.Println("Server is running on port ",server.conf.HttpPort)
	err := http.ListenAndServe(addr,wrappedMux)
	if err != nil {
		fmt.Println("Internal server error",err)
		os.Exit(1)
	}
	
}