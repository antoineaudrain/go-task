#!/bin/bash
export ENV="development"
export NATS_URL="localhost:4222"
export DATABASE_URL="postgresql://go_task:go_task@localhost:5432/user_db?sslmode=disable"
./user
