-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, userId, last_fetched_at)
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;
