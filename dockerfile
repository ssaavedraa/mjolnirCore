# Use the official Golang image with Go version 1.22.1 as a base image
FROM golang:1.22.1-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the /cmd directory into the container's /app directory
COPY ./cmd /app

# Build the Go app
RUN go build -o /app/out ./cmd

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/out /app/out

# Command to run the executable
CMD ["/app/out"]
