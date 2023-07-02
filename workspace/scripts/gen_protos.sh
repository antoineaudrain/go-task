#!/bin/bash

API_DIR="$(pwd)/api"

if ! protoc -I "${API_DIR}/pb" \
    --go_out="${API_DIR}" --go_opt=paths=source_relative \
    --go-grpc_out="${API_DIR}" --go-grpc_opt=paths=source_relative \
    "${API_DIR}/pb/workspace.proto"; then
    echo "Failed to generate workspace proto files"
    exit 1
fi

echo "Workspace proto files generated successfully!"
