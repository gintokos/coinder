# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /usr/src/app

# Install necessary build tools
RUN apk --no-cache add bash git make gcc libc-dev

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build the application
COPY . .
RUN make build OS=LINUX

# Final stage
FROM alpine

# Copy the build directory and config file
COPY --from=builder /usr/src/app/build ./build
COPY appconfig.json ./

# Make the binary executable
RUN chmod +x ./build/app/app

# Run the application with config
ENTRYPOINT ["./build/app/app", "--config", "./appconfig.json"]