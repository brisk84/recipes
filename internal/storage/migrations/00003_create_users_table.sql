-- +goose Up
CREATE TABLE users (
	login text NOT NULL,
	password text NOT NULL
);

-- +goose Down
DROP TABLE users;
