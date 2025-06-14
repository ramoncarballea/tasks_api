#!/bin/sh

# Load environment variables from .env if present
if [ -f .env ]; then
  # shellcheck disable=SC2046
  export $(grep -v '^#' .env | xargs)
fi

export MIGRATIONS_PATH=./config/database/migrations

goose -dir "$MIGRATIONS_PATH" postgres "$CONNECTION_STRING" up
