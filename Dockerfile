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
ARG KAFKA_BROKER

# Railway variable
ARG RAILWAY_PUBLIC_DOMAIN
ARG RAILWAY_PRIVATE_DOMAIN
ARG RAILWAY_PROJECT_NAME
ARG RAILWAY_ENVIRONMENT_NAME
ARG RAILWAY_SERVICE_NAME
ARG RAILWAY_PROJECT_ID
ARG RAILWAY_ENVIRONMENT_ID
ARG RAILWAY_SERVICE_ID

RUN apk add alpine-sdk
RUN apk --update add git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -tags musl -o /app/out ./cmd
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/out /app/out
CMD ["/app/out"]
