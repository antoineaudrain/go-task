#!/bin/bash
export ENV="production"
export NATS_URL="nats://nats.example.com:4222"
export DATABASE_URL="postgres://username:password@prod-db.example.com:5432/user_db"
./user
