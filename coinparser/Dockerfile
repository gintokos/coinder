FROM golang:1.24-alpine AS builder

WORKDIR /usr/src/app

RUN apk --no-cache add bash git make gcc libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o parser ./coinparser

FROM alpine

COPY --from=builder /usr/src/app/coinparser ./

COPY parserconfig.json ./

RUN chmod +x ./build/app/app ./build/parser/parser

RUN chmod +x ./coinparser

ENTRYPOINT ["./coinparser"]
