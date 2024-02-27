-- +goose Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(50) NOT NULL
);

-- +goose Down
DROP TABLE users;
