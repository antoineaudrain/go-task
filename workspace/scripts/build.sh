#!/bin/bash

protoc -I api/proto \
    --go_out=api --go_opt=paths=source_relative \
    --go-grpc_out=api --go-grpc_opt=paths=source_relative \
    api/proto/workspace.proto

go build -o auth-service ./cmd/main/main.go
