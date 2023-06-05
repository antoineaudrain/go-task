# Auth Service

This is the authentication microservice for a task management tool.

## Getting Started

### Prerequisites

- Go (1.17 or later)
- Docker (optional, for containerization)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/auth-service.git
    cd auth-service
    ```
   
2. Install dependencies:

    ```bash
    go mod download
    ```

### Usage

1. Generate gRPC Code:

    ```bash
   protoc -I api/proto \
    --go_out=api --go_opt=paths=source_relative \
    --go-grpc_out=api --go-grpc_opt=paths=source_relative \
    api/proto/workspace.proto
    ```

2. Build the service:

    ```bash
    go build -o auth-service ./cmd/main/main.go
    ```

3. Run the service:

    ```bash
    ./auth-service
    ```

By default, the service will listen on port 50051.

### Docker Support

To build and run the service in a Docker container, make sure you have Docker installed.

1. Build the Docker image:

    ```bash
    docker build -t auth-service .
    ```

2. Run the Docker container:

    ```bash
    docker run -p 50051:50051 auth-service
    ```

### API Documentation

The service follows the gRPC API. The API definition can be found in api/v1/auth_service.proto. You can generate client code from this file to interact with the service.

