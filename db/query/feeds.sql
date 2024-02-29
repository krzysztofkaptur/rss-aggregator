-- name: CreateFeed :one
INSERT INTO feeds ( name, url, user_id) VALUES($1, $2, $3) RETURNING *;

-- name: FetchFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds SET last_fetched_at=NOW(), updated_at=NOW() WHERE id=$1 RETURNING *;