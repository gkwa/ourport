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
WHERE  LOWER(url) NOT LIKE '%thumb%'
 AND   LOWER(url) NOT LIKE '%th-%'
 AND  (
   LOWER(url) LIKE '%.jpg'
   OR LOWER(url) LIKE '%.jpeg'
   OR LOWER(url) LIKE '%.gif'
   OR LOWER(url) LIKE '%.png'
 )
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