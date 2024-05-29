-- sqlc didn't generate properly cause Up in goose Up was in small case🤦.

-- +goose Up
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;