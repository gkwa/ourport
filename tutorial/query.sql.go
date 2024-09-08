package tutorial

import (
	"context"
	"database/sql"
)

const getImageLinks = `-- name: GetImageLinks :many
SELECT DISTINCT url,
               title,
               MIN(created_at) AS first_seen
FROM   links
WHERE  url NOT LIKE '%thumb%'
 AND  url NOT LIKE '%th-%'
GROUP  BY url
ORDER  BY first_seen
`

type GetImageLinksRow struct {
	Url       string
	Title     sql.NullString
	FirstSeen interface{}
}

func (q *Queries) GetImageLinks(ctx context.Context) ([]GetImageLinksRow, error) {
	rows, err := q.db.QueryContext(ctx, getImageLinks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetImageLinksRow
	for rows.Next() {
		var i GetImageLinksRow
		if err := rows.Scan(&i.Url, &i.Title, &i.FirstSeen); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
