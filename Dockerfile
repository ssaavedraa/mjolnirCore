# Use the official Golang image with Go version 1.22.1 as a base image
FROM golang:1.22.1-alpine AS build

# Specify the variable you need
# Service specific variables
ARG ICLOUD_AUTH_EMAIL
ARG ICLOUD_SENDER_EMAIL
ARG ICLOUD_PASSWORD
ARG SMTP_HOST
ARG SMTP_PORT
ARG ENVIRONMENT
ARG DOMAIN
ARG JWT_SECRET
ARG DB_HOST
ARG DB_NAME
ARG DB_PASSWORD
ARG DB_PORT
ARG DB_USER

# Railway variable
ARG RAILWAY_PUBLIC_DOMAIN
ARG RAILWAY_PRIVATE_DOMAIN
ARG RAILWAY_PROJECT_NAME
ARG RAILWAY_ENVIRONMENT_NAME
ARG RAILWAY_SERVICE_NAME
ARG RAILWAY_PROJECT_ID
ARG RAILWAY_ENVIRONMENT_ID
ARG RAILWAY_SERVICE_ID

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go app
RUN go build -o /app/out ./cmd

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=build /app/out /app/out

# Command to run the executable
CMD ["/app/out"]
