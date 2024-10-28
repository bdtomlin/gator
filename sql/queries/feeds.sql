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

-- name: MarkFeedFetched :exec
update feeds
set last_fetched_at = $2, updated_at = $2
where id = $1;

-- name: GetNextFeedToFetch :one
select * from feeds
order by last_fetched_at asc nulls first
limit 1;
