-- name: CreateFeed :one
insert into feeds (id, created_at, updated_at, name, url, user_id)
values ($1,$2,$3,$4,$5,$6)
returning *;

-- name: GetFeeds :many
select * from feeds;

-- name: GetFeedByURL :one
select * from feeds
where url = $1;

-- name: MarkFeedFetched :exec
update feeds
set updated_at = $1, last_fetched_at = $1
where id = $2;

-- name: GetNextFeedToFetch :one
select * from feeds
order by last_fetched_at asc nulls first;