package core

import (
	"context"
	"database/sql"
	"path/filepath"

	"github.com/go-logr/logr"

	_ "embed"

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

func Run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "links.sqlite")
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	queries := tutorial.New(db)

	links, err := queries.GetImageLinks(ctx)
	if err != nil {
		return err
	}

	groups := make(map[string][]string)
	for _, link := range links {
		parentDir := filepath.Dir(link.Url)
		groups[parentDir] = append(groups[parentDir], link.Url)
	}

	return nil
}
