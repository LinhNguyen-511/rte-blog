package data

import (
	"database/sql"
	"log"
	"os"
	"rte-blog/types"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const defaultDbName = "postgres"
const testDbName = "rte_blog_test"

var db *sql.DB

var dbConfig = types.DbConfig{
	DbName:   defaultDbName,
	User:     "tester",
	Password: "testing123",
}

var dbConfigMigration = types.DbConfig{
	DbName:   testDbName,
	User:     "tester",
	Password: "testing123",
}

func TestMain(m *testing.M) {
	db = getTestDatabaseConnection()

	err := createTestDatabase(db, testDbName)
	defer db.Close()

	if err != nil {
		log.Print(err)
	}
	runMigrations()

	exitVal := m.Run()

	deleteTestDatabase(db, testDbName)

	os.Exit(exitVal)
}

func getTestDatabaseConnection() *sql.DB {
	return Connect(dbConfig)
}

func createTestDatabase(db *sql.DB, testDbName string) error {
	_, err := db.Exec("CREATE DATABASE " + testDbName)
	return err
}

func runMigrations() {
	connectionString := generateConnectionString(dbConfigMigration)
	migration, err := migrate.New(
		"file://migrations",
		connectionString,
	)

	if err != nil {
		log.Print(err)
	}

	if err := migration.Up(); err != nil {
		log.Print(err)
	}

	migration.Close()
}

func deleteTestDatabase(db *sql.DB, testDbName string) error {
	_, err := db.Exec("DROP DATABASE " + testDbName)

	if err != nil {
		log.Print(err)
	}

	return err
}
