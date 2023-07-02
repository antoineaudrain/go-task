#!/bin/bash

API_DIR="$(pwd)/api"

if ! protoc -I "${API_DIR}/pb" \
    --go_out="${API_DIR}" --go_opt=paths=source_relative \
    --go-grpc_out="${API_DIR}" --go-grpc_opt=paths=source_relative \
    "${API_DIR}/pb/invitation.proto"; then
    echo "Failed to generate invitation proto files"
    exit 1
fi

echo "Invitation proto files generated successfully!"
