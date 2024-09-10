# Build the Go API
FROM golang:1.23.1-alpine3.20

WORKDIR /usr/src/app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the application source code
COPY . .

# Build the API binary
RUN go build -o /usr/local/bin/app ./cmd/api/main.go

EXPOSE 5000

# Run the API
CMD app

HEALTHCHECK CMD curl --fail http://localhost:5000/healthcheck || exit 1
