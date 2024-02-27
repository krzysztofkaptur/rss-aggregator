-- +goose Up
-- add unique for email
-- add created_at and updated_at
-- create apikey
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(50) NOT NULL
);

-- +goose Down
DROP TABLE users;
