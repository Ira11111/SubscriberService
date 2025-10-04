-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- индекс для поиска по id
CREATE INDEX idx_subscriptions_user_id ON subscription_user(id_user);
CREATE INDEX idx_subscriptions_id_user ON subscription_user(id_sub);
CREATE INDEX idx_subscriptions_id ON subscription(id);

-- индекс для поиска оп имени сервиса
CREATE INDEX idx_subscriptions_name ON subscription(name)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP INDEX idx_subscriptions_user_id;
DROP INDEX idx_subscriptions_id_user;
DROP INDEX idx_subscriptions_id;
DROP INDEX idx_subscriptions_name;
-- +goose StatementEnd
