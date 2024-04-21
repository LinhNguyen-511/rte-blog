package server

import (
	"log"
	"net/http"
	"rte-blog/data"

	"github.com/labstack/echo/v4"
)

type server struct {
	config    *http.Server
	postModel *data.PostModel
}

// TODO: maybe create an interface for server config

func New() *server {
	store, err := data.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return &server{
		config: &http.Server{
			Addr: ":3000",
		},
		postModel: &data.PostModel{Store: store},
	}
}

func (server *server) Init() {
	echo := echo.New()
	server.config.Handler = echo

	echo.GET("/", home)
	echo.POST("/posts", server.handleCreatePost)
	echo.GET("/posts/:id", server.handleGetPost)

	// TODO: change to HTTPS server before release: https://echo.labstack.com/docs/start-server#https-server
	if err := server.config.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
