#!/bin/bash
export ENV="production"
export SECRET_KEY="Q85yAn4EDHBzaCtHERTChGaFej84J5GFYQPo3fdC"
export NATS_URL="nats://nats.example.com:4222"
export DATABASE_URL="postgres://username:password@prod-db.example.com:5432/workspace_db"
export GOOSE_DRIVER="postgres"
export GOOSE_DBSTRING="postgresql://go_task:go_task@localhost:5432/workspace_db?sslmode=disable"
./workspace
