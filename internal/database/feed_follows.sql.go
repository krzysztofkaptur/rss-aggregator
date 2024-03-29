// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: feed_follows.sql

package database

import (
	"context"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feed_follows (user_id, feed_id) VALUES ($1, $2) RETURNING id, created_at, updated_at, user_id, feed_id
`

type CreateFeedFollowParams struct {
	UserID int32 `json:"user_id"`
	FeedID int32 `json:"feed_id"`
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow, arg.UserID, arg.FeedID)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const deleteFeedFollows = `-- name: DeleteFeedFollows :one
DELETE FROM feed_follows WHERE user_id=$1 AND feed_id=$2 RETURNING id, created_at, updated_at, user_id, feed_id
`

type DeleteFeedFollowsParams struct {
	UserID int32 `json:"user_id"`
	FeedID int32 `json:"feed_id"`
}

func (q *Queries) DeleteFeedFollows(ctx context.Context, arg DeleteFeedFollowsParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, deleteFeedFollows, arg.UserID, arg.FeedID)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const fetchFeedFollows = `-- name: FetchFeedFollows :many
SELECT id, created_at, updated_at, user_id, feed_id FROM feed_follows WHERE user_id=$1
`

func (q *Queries) FetchFeedFollows(ctx context.Context, userID int32) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, fetchFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
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
