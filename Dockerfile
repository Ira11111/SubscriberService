FROM golang:1.24.0-alpine3.20 AS builder
WORKDIR /usr/local/src
RUN apk --no-cache add bash make gcc gettext git musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY ./api ./api
COPY ./internal ./internal
COPY ./cmd ./cmd
RUN go build -o ./bin/app ./cmd/main.go

FROM alpine
RUN apk --no-cache add bash
WORKDIR /usr/local/src
COPY --from=builder /usr/local/src/bin/app ./
EXPOSE 8080
