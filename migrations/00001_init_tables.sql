-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE subscription(
    id serial primary key,
    name varchar(50) not null,
    price integer default 0
);
CREATE TABLE subscription_user(
    id serial primary key,
    id_sub integer,
    id_user varchar(32),
    start_date date default NOW(),
    end_date date default null,
    foreign key (id_sub) references subscription on delete no action
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE subscription_user;
DROP TABLE subscription;
-- +goose StatementEnd
