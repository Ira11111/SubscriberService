-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE subscription_user
DROP CONSTRAINT subscription_user_id_sub_fkey;


ALTER TABLE subscription_user
ADD CONSTRAINT subscription_user_id_sub_fkey
FOREIGN KEY (id_sub) REFERENCES subscription(id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE subscription_user
DROP CONSTRAINT subscription_user_id_sub_fkey;

ALTER TABLE subscription_user
ADD CONSTRAINT subscription_user_id_sub_fkey
FOREIGN KEY (id_sub) REFERENCES subscription(id)
ON DELETE RESTRICT;
-- +goose StatementEnd
