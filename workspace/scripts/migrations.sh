#!/bin/bash
MIGRATIONS_DIR="./migrations"
goose -dir $MIGRATIONS_DIR postgres "$DATABASE_URL" "$1"
