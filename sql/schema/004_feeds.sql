-- +goose UP
alter table feeds
add column last_fetched_at timestamp;

-- +goose DOWN
alter table feeds
remove column last_fetched_at;