-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
VALUES ( $1, $2, $3, $4, $5 )
RETURNING *,
(select name as user_name from users u where u.id = user_id),
(select name as feed_name from feeds f where f.id = feed_id);

-- name: GetFeedFollowsForUser :many
select feed_follows.*, users.name as user_name, feeds.name as feed_name
from feed_follows 
inner join users on users.id = feed_follows.user_id
inner join feeds on feeds.id = feed_follows.feed_id
where feed_follows.user_id = $1;
