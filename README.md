# Gator RSS Cli

## Requirements

- Go programming language
- Postgresql

## Installation

If you have go installed on your system then use go install like this:

```
go install https://github.com/bdtomlin/gator

```

## Usage

create a .gatorconfig.json file in your home directory with the following:

```json
{
  "db_url": "postgres://postgres:@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

The following commands are available:

| command  | args     | description      | example                      |
| -------- | -------- | ---------------- | ---------------------------- |
| login    | username | logs user in     | `gator login jimbo`          |
| register | username | creates new user | `gator register jimbo`       |
| reset    | (none)   | deletes data     | `gator reset`                |
| users    | (none)   | lists the users  | `gator users`                |
| agg      | delay    | saves feeds      | `gator agg 3s`               |
| addfeed  | url      | adds a feed      | `gator addfeed https://...`  |
| feeds    | ulr      | list feeds       | `gator feeds`                |
| follow   | url      | follow a feed    | `gator follow https://...`   |
| unfollow | url      | unfollow a feed  | `gator unfollow https://...` |
| browse   | limit    | browse feeds     | `gator browse 5`             |
