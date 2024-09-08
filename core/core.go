package core

import (
	"github.com/go-logr/logr"

	"context"
	"database/sql"
	_ "embed"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gkwa/ourport/tutorial"
)

func Hello(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Hello function")
	logger.Info("Hello, World!")
	logger.V(1).Info("Debug: Exiting Hello function")
}

//go:embed schema.sql
var ddl string

func run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "links.sqlite")
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	queries := tutorial.New(db)

	// list all authors
	links, err := queries.GetLinks(ctx)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
