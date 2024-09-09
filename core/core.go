package core

import (
	"context"
	"database/sql"
	"path"
	"strconv"
	"strings"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gkwa/ourport/tutorial"
)

//go:embed schema.sql
var ddl string

func FetchImageLinks() ([]tutorial.GetImageLinksRow, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "links.sqlite")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	queries := tutorial.New(db)
	return queries.GetImageLinks(ctx)
}

func extractNumber(url string) int {
	base := path.Base(url)
	numStr := strings.TrimSuffix(base, path.Ext(base))
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0
	}
	return num
}
