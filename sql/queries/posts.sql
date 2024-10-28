-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
select p.* from posts p
inner join feeds f on f.id = p.feed_id
where f.user_id = $1
order by p.published_at desc
limit $2;

-- name: PostWithUrlExists :one
select EXISTS(select 1 from posts where url = $1);
