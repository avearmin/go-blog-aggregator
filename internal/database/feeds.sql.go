// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: feeds.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createFeed = `-- name: createFeed :one
INSERT INTO feeds (name, url, userId)
VALUES ($1, $2, $3) 
RETURNING name, url, userid
`

type createFeedParams struct {
	Name   string
	Url    string
	Userid uuid.NullUUID
}

func (q *Queries) createFeed(ctx context.Context, arg createFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed, arg.Name, arg.Url, arg.Userid)
	var i Feed
	err := row.Scan(&i.Name, &i.Url, &i.Userid)
	return i, err
}