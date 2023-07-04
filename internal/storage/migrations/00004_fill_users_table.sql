-- +goose Up
INSERT INTO users (login,"password") VALUES
	 ('user01','$2a$10$IjY1vdDM41oSNcoPwN8ipOlt3bg0J4cn2afz5RJ8RxEI5CglF7/V.');

-- +goose Down
DELETE from users;