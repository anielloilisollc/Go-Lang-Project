# Use Golang 1.27 for building the app
FROM golang:1.27-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the application source code to the container
COPY . .

# Build the Go app as an executable named 'main'
RUN go build -o main .

# Use a smaller image to run the app
FROM alpine:latest

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /main

# Expose the port the app will listen on
EXPOSE 8080

# Run the application
CMD ["/main"]
