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

# Create necessary directories
RUN mkdir -p ./build/app ./build/parser

# Copy the binaries and config files
COPY --from=builder /usr/src/app/build/app/app ./build/app/
COPY --from=builder /usr/src/app/build/parser/parser ./build/parser/
COPY appconfig.json ./
COPY parserconfig.json ./
COPY photo ./photo

# Make the binaries executable
RUN chmod +x ./build/app/app ./build/parser/parser

# Create script to run both applications
RUN echo '#!/bin/sh' > ./run.sh && \
    echo './build/app/app --config ./appconfig.json &' >> ./run.sh && \
    echo './build/parser/parser --config ./parserconfig.json &' >> ./run.sh && \
    echo 'wait' >> ./run.sh && \
    chmod +x ./run.sh

# Run the applications with config
ENTRYPOINT ["./run.sh"]