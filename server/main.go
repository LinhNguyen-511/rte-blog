package server

import (
	"database/sql"
	"log"
	"net/http"
	"rte-blog/data"

	"github.com/labstack/echo/v4"
)

type model struct {
	db *sql.DB
}

func Init() {
	echo := echo.New()

	db, err := data.Connect()
	if err != nil {
		log.Fatal(err)
	}

	model := &model{db}

	echo.GET("/", model.GetPost)

	server := http.Server{
		Addr:    ":3000",
		Handler: echo,
	}

	// TODO: change to HTTPS server before release: https://echo.labstack.com/docs/start-server#https-server
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
