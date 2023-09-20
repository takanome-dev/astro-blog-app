// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: posts.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, title, body, user_id, is_published, is_draft) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, title, body, user_id, is_published, is_draft, created_at, updated_at, deleted_at
`

type CreatePostParams struct {
	ID          uuid.UUID
	Title       string
	Body        string
	UserID      uuid.UUID
	IsPublished bool
	IsDraft     bool
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.Title,
		arg.Body,
		arg.UserID,
		arg.IsPublished,
		arg.IsDraft,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.IsPublished,
		&i.IsDraft,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getAllPosts = `-- name: GetAllPosts :many
SELECT id, title, body, user_id, is_published, is_draft, created_at, updated_at, deleted_at FROM posts
`

func (q *Queries) GetAllPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.UserID,
			&i.IsPublished,
			&i.IsDraft,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostByID = `-- name: GetPostByID :one
SELECT id, title, body, user_id, is_published, is_draft, created_at, updated_at, deleted_at FROM posts WHERE id = $1
`

func (q *Queries) GetPostByID(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.IsPublished,
		&i.IsDraft,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
