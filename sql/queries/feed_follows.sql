-- name: CreateFeedFollow :one
with insert_feed_follow as (
    insert into feed_follows (id, created_at, updated_at, user_id, feed_id)
    values ($1,$2,$3,$4,$5)
    returning *
)
select
    insert_feed_follow.*,
    feeds.name as feed_name,
    users.name as user_name
from insert_feed_follow
inner join feeds on insert_feed_follow.feed_id = feeds.id
inner join users on insert_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
select
    feed_follows.*,
    users.name as user_name,
    feeds.name as feed_name
from feed_follows
inner join feeds on feed_follows.feed_id = feeds.id
inner join users on feed_follows.user_id = users.id
where feed_follows.user_id = $1;