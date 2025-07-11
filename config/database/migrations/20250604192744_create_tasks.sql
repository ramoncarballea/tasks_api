-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS tasks(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    expires_at timestamp NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp);

CREATE INDEX idx_tasks_name ON tasks(name);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP INDEX IF EXISTS idx_tasks_name;

DROP TABLE IF EXISTS tasks;