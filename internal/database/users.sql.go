// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email) VALUES ($1) RETURNING id, email, created_at, updated_at, api_key
`

func (q *Queries) CreateUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ApiKey,
	)
	return i, err
}

const fetchUsers = `-- name: FetchUsers :many
SELECT id, email, created_at, updated_at, api_key FROM users
`

func (q *Queries) FetchUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, fetchUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ApiKey,
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

const getUserByAPIKey = `-- name: GetUserByAPIKey :one
SELECT id, email, created_at, updated_at, api_key FROM users WHERE api_key=$1
`

func (q *Queries) GetUserByAPIKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAPIKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ApiKey,
	)
	return i, err
}
