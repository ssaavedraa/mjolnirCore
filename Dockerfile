# Use the official Golang image with Go version 1.22.1 as a base image
FROM golang:1.22.1-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Copy the .env file into the container
COPY .env .

# Build the Go app
RUN go build -o /app/out ./cmd

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=build /app/out /app/out

# Copy the .env file into the container
COPY --from=build /app/.env /app/.env

# Command to run the executable
CMD ["/app/out"]
