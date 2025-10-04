-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE INDEX idx_subscriptions_user ON subscription_user(id_sub, id_user);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP INDEX idx_subscriptions_user;
-- +goose StatementEnd
