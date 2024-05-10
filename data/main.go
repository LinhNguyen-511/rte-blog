package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"rte-blog/types"
)

func GenerateConnectionString(config types.DbConfig) string {
	return fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", config.DbName, config.User, config.Password)
}

func Connect(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
