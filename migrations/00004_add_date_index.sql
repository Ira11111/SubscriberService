-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
--индексы для диапазонов дат
CREATE INDEX idx_sub_start_date ON subscription_user(start_date);
CREATE INDEX idx_sub_end_date ON subscription_user(end_date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP INDEX idx_sub_start_date;
DROP INDEX idx_sub_end_date;
-- +goose StatementEnd
