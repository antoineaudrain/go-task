# Use the official Go image from the Docker Hub
FROM golang:1.17

# Enable Go modules
ENV GO111MODULE=on

# Set the current working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum file to download the dependencies
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

# This container exposes port 50051 to the outside world
EXPOSE 50051

# Run the executable
CMD ["./main"]
