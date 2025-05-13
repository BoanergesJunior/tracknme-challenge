FROM golang:1.23.3-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code and migrations
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/application

# Expose the port
EXPOSE ${PORT}

# Command to run the executable
CMD ["./main"]
