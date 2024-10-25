-- +goose Up
-- +goose StatementBegin
create table users(
  id uuid primary key,
  created_at timestamp not null,
  updated_at timestamp not null,
  name text unique not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
