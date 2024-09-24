# Step 1: Build the Go binary
FROM golang:1.23.1-alpine AS build

# Step 2: setup environment
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Step 3: set up the working directory
WORKDIR /app/disester-management

# Step 4:copy sum and mod file to the working directory
COPY go.mod go.sum ./

# Step 5:run go get the packages
RUN go mod download

# Step 6:copy the rest of the application source files to the working directory
COPY . .

# Build the Go binary from the cmd/server/main.go path
RUN go build -o /app/bin/main ./cmd/server/main.go

# Step 2: Create a smaller container to run the Go binary
FROM alpine:latest

# Set the working directory inside the new container
WORKDIR /app


# Copy the binary from the builder container
COPY --from=build /app/bin/main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]


