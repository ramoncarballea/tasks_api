-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS roles(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp
);

CREATE UNIQUE INDEX idx_roles_name ON roles(name);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP INDEX IF EXISTS idx_roles_name;

DROP TABLE IF EXISTS roles;
