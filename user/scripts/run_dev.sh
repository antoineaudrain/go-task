#!/bin/bash
export ENV="development"
export SECRET_KEY="hn9yCT45BP4dTj5TFJK9E8ojm98Jdfkp4NLF3rdd"
export NATS_URL="localhost:4222"
export DATABASE_URL="postgresql://go_task:go_task@localhost:5432/user_db?sslmode=disable"
export GOOSE_DRIVER="postgres"
export GOOSE_DBSTRING="postgresql://go_task:go_task@localhost:5432/user_db?sslmode=disable"
./user
