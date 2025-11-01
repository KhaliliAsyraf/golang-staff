# Use official Golang image
FROM golang:1.23.2

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the project
COPY . .

CMD ["go", "run", ".", "--port", "8080"]