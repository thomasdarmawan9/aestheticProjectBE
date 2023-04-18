# Base Image
FROM golang:1.19

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code from the host to the container
COPY go.mod .

# Build the Go app
RUN go mod tidy

COPY . .

RUN go build -o server ./cmd/server

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./server"]
