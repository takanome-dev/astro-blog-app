-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY NOT NULL,
    title TEXT NOT NULL UNIQUE DEFAULT '',
    body TEXT NOT NULL DEFAULT '',
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    is_published BOOLEAN NOT NULL DEFAULT FALSE,
    is_draft BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
DROP TABLE posts;