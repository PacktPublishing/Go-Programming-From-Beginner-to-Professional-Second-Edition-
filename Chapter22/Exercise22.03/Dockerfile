FROM golang:latest AS builder

# Ensure Go compiler builds a statically-linked binary, including all necessary libraries within the binary.
ENV CGO_ENABLED=0

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
COPY main.go ./
RUN go mod download

RUN go build -o monitored_app .

# Start a new stage to create a minimal image
FROM scratch

# Copy the compiled binary from the builder stage
COPY --from=builder /app/monitored_app /.

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./monitored_app"]
