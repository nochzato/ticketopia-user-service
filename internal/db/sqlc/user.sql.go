// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
	username,
	full_name,
	password,
	email
) VALUES (
	$1, $2, $3, $4
)
RETURNING id, full_name, username, password, email, created_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.FullName,
		arg.Password,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, full_name, username, password, email, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
	password = COALESCE($1, password),
	full_name = COALESCE($2, full_name),
	email = COALESCE($3, email)
WHERE
    id = $4
RETURNING id, full_name, username, password, email, created_at
`

type UpdateUserParams struct {
	Password pgtype.Text `json:"password"`
	FullName pgtype.Text `json:"full_name"`
	Email    pgtype.Text `json:"email"`
	ID       uuid.UUID   `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Password,
		arg.FullName,
		arg.Email,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}