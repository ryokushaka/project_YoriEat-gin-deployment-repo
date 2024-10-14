FROM public.ecr.aws/docker/library/golang:1.22-alpine AS builder

# Install necessary packages
RUN apk add --no-cache git bash

# Create working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download and tidy Go modules
RUN go mod download
RUN go mod tidy

# Copy the rest of the files
COPY . .

# Install Swagger CLI and generate documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Execution stage
FROM public.ecr.aws/docker/library/alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /root/

# Copy the built files and necessary files
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Command to run the container
CMD ["./main"]