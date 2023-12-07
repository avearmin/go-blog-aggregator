-- name: createFeed :one
INSERT INTO feeds (name, url, userId)
VALUES ($1, $2, $3) 
RETURNING *;
