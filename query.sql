-- name: GetLinks :many
SELECT DISTINCT url,
                title,
                MIN(created_at) AS first_seen
FROM   links
GROUP  BY url
ORDER  BY first_seen; 

-- name: CreateLink :one
INSERT
or     IGNORE
into   links
       (
              url,
              title
       )
       VALUES
       (
              ?,
              ?
       )
RETURNING *;