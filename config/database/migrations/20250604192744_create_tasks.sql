CREATE TABLE IF NOT EXISTS tasks(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    expires_at timestamp NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp);