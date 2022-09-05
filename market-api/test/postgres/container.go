package postgres

import (
	"database/sql"
	"fmt"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"log"
	"path/filepath"
	"runtime"
)

func SetUp() (*gnomock.Container, *sql.DB, error) {
	const user = "test"
	const pass = "test"
	const dbName = "test"

	_, path, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("failed to get path")
	}

	filepath := filepath.Dir(path) + "/init.sql"

	p := postgres.Preset(
		postgres.WithUser(user, pass),
		postgres.WithDatabase(dbName),
		postgres.WithQueriesFile(filepath),
	)

	container, err := gnomock.Start(p)
	if err != nil {
		log.Fatal(err.Error())
		return nil, nil, err
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
		container.Host, container.DefaultPort(), user, pass, dbName,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return nil, nil, err
	}

	return container, db, nil
}
