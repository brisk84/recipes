-- +goose Up
CREATE TABLE if not exists recipes (
	id text NOT NULL,
	cr_dt timestamptz NULL,
	upd_dt timestamptz NULL,
	del_dt timestamptz NULL,
	user_id text NULL,
	title text NOT NULL,
	description text NULL,
	ingredients text[] NULL,
	steps text[] NULL,
	CONSTRAINT recipes_pkey PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE recipes;
