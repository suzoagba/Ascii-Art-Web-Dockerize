# Use the official Golang image as the base image
FROM golang:alpine

LABEL name="ascii-art-web-dockerize"

LABEL description="A simple web server that displays ASCII art banners"

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules and download them
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Use alpine as the base image for the final image
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the executable from the builder image to the final image
COPY --from=0 /app/main .

# Expose the port on which the application will run
EXPOSE 8080

# Start the application
CMD ["./main"]
