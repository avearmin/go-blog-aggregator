-- +goose Up

CREATE TABLE feeds (
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    userId UUID REFERENCES users(id) ON DELETE CASCADE
);
