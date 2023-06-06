#!/bin/bash
export ENV="production"
export SECRET_KEY="Q85yAn4EDHBzaCtHERTChGaFej84J5GFYQPo3fdC"
export NATS_URL="nats://nats.example.com:4222"
export DATABASE_URL="postgres://username:password@prod-db.example.com:5432/user_db"
./user
