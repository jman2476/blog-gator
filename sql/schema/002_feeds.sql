-- +goose UP
CREATE TABLE feeds(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE,
    url TEXT UNIQUE,
    user_id UUID NOT NULL
                REFERENCES users(id)
                ON DELETE CASCADE
);

-- +goose DOWN
DROP TABLE feeds;