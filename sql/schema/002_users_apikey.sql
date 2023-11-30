-- +goose Up
ALTER TABLE users
ADD COLUMN apikey VARCHAR(64) UNIQUE NOT NULl
DEFAULT encode(sha256(random()::text::bytea), 'hex');