#!/bin/bash
export ENV="development"
export SECRET_KEY="hn9yCT45BP4dTj5TFJK9E8ojm98Jdfkp4NLF3rdd"
export NATS_URL="localhost:4222"
export DATABASE_URL="postgresql://go_task:go_task@localhost:5433/workspace_db?sslmode=disable"
./workspace
