-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
  id uuid primary key,
  created_at timestamp not null,
  updated_at timestamp not null,
  title text not null,
  url text unique not null,
  description text not null,
  published_at timestamp not null,
  feed_id uuid not null,
  FOREIGN KEY(feed_id) REFERENCES feeds(id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
