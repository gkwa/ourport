package core

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gkwa/ourport/tutorial"
)

//go:embed schema.sql
var ddl string

func Run() error {
	// Existing Run function code...
	return nil
}

func Report1() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "links.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := tutorial.New(db)

	links, err := fetchImageLinks(ctx, queries)
	if err != nil {
		return err
	}

	generateReport1(links)

	return nil
}

func fetchImageLinks(ctx context.Context, queries *tutorial.Queries) ([]tutorial.GetImageLinksRow, error) {
	return queries.GetImageLinks(ctx)
}

func generateReport1(links []tutorial.GetImageLinksRow) {
	groups := make(map[string][]string)
	for _, link := range links {
		parentDir := filepath.Dir(link.Url)
		groups[parentDir] = append(groups[parentDir], link.Url)
	}

	fmt.Printf("Number of groups: %d\n", len(groups))
	fmt.Printf("Total number of links: %d\n", len(links))
}
