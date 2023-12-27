# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY ./my_app /app


# Download and install any required dependencies
RUN go mod download

# Build the Go app

# Expose port 8080 for incoming traffic
EXPOSE 8080


CMD ["go","run","main.go"]