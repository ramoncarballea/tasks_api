@echo off

REM Set DATABASE_URL here or ensure it is set in your environment

set CONNECTION_STRING=postgres://postgres:postgres@localhost:5432/tasks_db?sslmode=disable
set MIGRATIONS_PATH=./config/database/migrations postgres

goose -dir %MIGRATIONS_PATH% "%CONNECTION_STRING%" up
