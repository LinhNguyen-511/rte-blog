package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"rte-blog/types"
)

func generateConnectionString(config types.DbConfig) string {
	return fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", config.User, config.Password, config.DbName)
}

func Connect(config types.DbConfig) *sql.DB {
	db, err := sql.Open("postgres", generateConnectionString(config))

	if err != nil {
		log.Fatal(err)
	}

	return db
}
