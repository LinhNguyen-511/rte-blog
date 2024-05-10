package server

import (
	"log"
	"net/http"
	"os"
	"rte-blog/data"
	"rte-blog/types"

	"github.com/labstack/echo/v4"
)

type server struct {
	config    *http.Server
	postModel data.PostStore
}

// TODO: maybe create an interface for server config

func New() *server {
	dbConfig := types.DbConfig{
		DbName:   "rte_blog",
		User:     os.Getenv("PQ_USERNAME"),
		Password: os.Getenv("PQ_PASSWORD"),
	}

	store := data.Connect(data.GenerateConnectionString(dbConfig))
	defer store.Close()

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

	// echo.GET("/", )
	echo.POST("/posts", server.handleCreatePost)
	echo.GET("/posts/:id", server.handleGetPost)
	echo.PUT("/posts/:id/title", server.handlePutPostTitle)

	// TODO: change to HTTPS server before release: https://echo.labstack.com/docs/start-server#https-server
	if err := server.config.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
