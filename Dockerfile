# Use an official Golang image as the base image
FROM golang:1.23-alpine as builder

# Install git and other dependencies
RUN apk add --no-cache git curl

# Set the Current Working Directory inside the container
WORKDIR /app

# Install air for live reloading from the new repository location
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Expose the port that the server will listen on
EXPOSE 8080

# Command to run the app with Air for live reloading
CMD ["air", "-c", ".air.toml"]
