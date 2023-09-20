// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID
	Title       string
	Body        string
	UserID      uuid.UUID
	IsPublished bool
	IsDraft     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
