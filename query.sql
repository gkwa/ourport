-- name: GetLinks :many
SELECT DISTINCT url,
                title,
                MIN(created_at) AS first_seen
FROM   links
GROUP  BY url
ORDER  BY first_seen; 

-- name: GetImageLinks :many
SELECT DISTINCT url,
                title,
                MIN(created_at) AS first_seen
FROM   links
WHERE  url NOT LIKE '%thumb%'
  AND  url NOT LIKE '%th-%'
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