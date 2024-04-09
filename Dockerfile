# Use the official Golang image with Go version 1.22.1 as a base image
FROM golang:1.22.1-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the entire project directory into the container
COPY . .

# Set the Current Working Directory to the cmd directory
WORKDIR /app/cmd

# Build the Go app
RUN go build -o /app/out .

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/out /app/out

# Command to run the executable
CMD ["/app/out"]
