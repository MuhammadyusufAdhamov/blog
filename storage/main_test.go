package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var (
	dbManager *DBManager

	//postgres credentials
	PostgresUser     = "postgres"
	PostgresPassword = "7"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresDatabase = "golang"
)

func TestMain(m *testing.M) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPassword,
		PostgresDatabase,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	dbManager = NewDBManager(db)
	os.Exit(m.Run())
}
