-- +goose Up
CREATE TABLE votes (
	cr_dt timestamptz NULL,
	recipe_id text NOT NULL,
	user_id text NOT NULL,
	mark integer NOT NULL,
	UNIQUE (recipe_id, user_id)
);

-- +goose Down
DROP TABLE votes;
