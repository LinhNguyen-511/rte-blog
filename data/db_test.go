package data

import (
	"database/sql"
	"log"
	"os"
	"rte-blog/types"
	"testing"

	"github.com/golang-migrate/migrate/v4"
)

func TestMain(m *testing.M) {
	const testDbName = "rte_blog_test"
	dbConfig := types.DbConfig{
		DbName:   "postgres",
		User:     os.Getenv("PQ_USERNAME"),
		Password: os.Getenv("PQ_PASSWORD"),
	}
	dbConnectionString := GenerateConnectionString(dbConfig)

	db := getTestDatabaseConnection(dbConnectionString)
	defer db.Close()

	createTestDatabase(db, testDbName)

	runMigrations(dbConnectionString)

	exitVal := m.Run()

	deleteTestDatabase(db, testDbName)

	os.Exit(exitVal)
}

func getTestDatabaseConnection(connectionString string) *sql.DB {
	db := Connect(connectionString)

	return db
}

func createTestDatabase(db *sql.DB, testDbName string) {
	_, err := db.Exec("CREATE DATABASE $1 IF NOT EXISTS", testDbName)

	if err != nil {
		log.Fatal(err)
	}
}

func runMigrations(connectionString string) {
	migration, err := migrate.New(
		"file://data/migrations",
		connectionString)

	if err != nil {
		log.Fatal(err)
	}

	if err := migration.Up(); err != nil {
		log.Fatal(err)
	}
}

func deleteTestDatabase(db *sql.DB, testDbName string) error {
	_, err := db.Exec("DROP DATABASE $1 IF EXISTS", testDbName)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
