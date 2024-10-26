-- +goose Up
-- +goose StatementBegin
create table feeds(
  id uuid primary key,
  created_at timestamp not null,
  updated_at timestamp not null,
  name text not null,
  url text not null unique,
  user_id uuid not null,
  FOREIGN KEY(user_id)
  REFERENCES users(id)
  on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table feeds;
-- +goose StatementEnd
