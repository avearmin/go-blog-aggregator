-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, userId, last_fetched_at)
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at IS NULL, last_fetched_at
LIMIT $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;
