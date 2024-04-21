package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	var (
		dbName   = "rte_blog"
		user     = os.Getenv("PQ_USERNAME")
		password = os.Getenv("PQ_PASSWORD")
	)

	connectionString := fmt.Sprintf("dbname=%s user=%s password=%s", dbName, user, password)
	db, err := sql.Open("postgres", connectionString)

	return db, err
}
