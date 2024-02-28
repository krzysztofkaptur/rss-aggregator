// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
)

type User struct {
	ID        int32        `json:"id"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	ApiKey    string       `json:"api_key"`
}
