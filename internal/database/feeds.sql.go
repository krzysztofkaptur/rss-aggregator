// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: feeds.sql

package database

import (
	"context"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds ( name, url, user_id) VALUES($1, $2, $3) RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at
`

type CreateFeedParams struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	UserID int32  `json:"user_id"`
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed, arg.Name, arg.Url, arg.UserID)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const fetchFeeds = `-- name: FetchFeeds :many
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at FROM feeds
`

func (q *Queries) FetchFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, fetchFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.LastFetchedAt,
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

const getNextFeedToFetch = `-- name: GetNextFeedToFetch :one
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1
`

func (q *Queries) GetNextFeedToFetch(ctx context.Context) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getNextFeedToFetch)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const markFeedAsFetched = `-- name: MarkFeedAsFetched :one
UPDATE feeds SET last_fetched_at=NOW(), updated_at=NOW() WHERE id=$1 RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at
`

func (q *Queries) MarkFeedAsFetched(ctx context.Context, id int32) (Feed, error) {
	row := q.db.QueryRowContext(ctx, markFeedAsFetched, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}
