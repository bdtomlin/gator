-- +goose Up
-- +goose StatementBegin
ALTER TABLE feeds
add column last_fetched_at timestamp;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE feeds
drop column last_fetched_at;
-- +goose StatementEnd
