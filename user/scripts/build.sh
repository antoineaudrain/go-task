#!/bin/bash

protoc -I api \
    --go_out=pkg/pb --go_opt=paths=source_relative \
    --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
    api/user.proto

go build -o user-service ./cmd/main/main.go
