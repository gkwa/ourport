// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package tutorial

import (
	"database/sql"
)

type Link struct {
	ID        int64
	Url       string
	Title     sql.NullString
	CreatedAt sql.NullTime
}
