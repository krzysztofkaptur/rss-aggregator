-- name: CreateFeedFollow :one
INSERT INTO feed_follows (user_id, feed_id) VALUES ($1, $2) RETURNING *;

-- name: FetchFeedFollows :many
SELECT * FROM feed_follows WHERE user_id=$1;

-- name: DeleteFeedFollows :one
DELETE FROM feed_follows WHERE user_id=$1 AND feed_id=$2 RETURNING *;