-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ( $1, $2, $3, $4, $5, $6 )
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsWithUsers :many
SELECT * FROM feeds inner join users
on users.id = feeds.user_id;

-- name: GetFeedForUrl :one
SELECT * FROM feeds where url = $1;
