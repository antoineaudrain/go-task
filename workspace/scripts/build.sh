#!/bin/bash

protoc -I api/v1 \
    --go_out=pkg/pb --go_opt=paths=source_relative \
    --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
    api/v1/auth_service.proto

go build -o auth-service ./cmd/main/main.go
