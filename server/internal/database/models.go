// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID    `json:"id"`
	Body      string       `json:"body"`
	UserID    uuid.UUID    `json:"user_id"`
	PostID    uuid.UUID    `json:"post_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type Post struct {
	ID          uuid.UUID    `json:"id"`
	Title       string       `json:"title"`
	Body        string       `json:"body"`
	UserID      uuid.UUID    `json:"user_id"`
	IsPublished bool         `json:"is_published"`
	IsDraft     bool         `json:"is_draft"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	Image       string       `json:"image"`
}

type User struct {
	ID        uuid.UUID    `json:"id"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
