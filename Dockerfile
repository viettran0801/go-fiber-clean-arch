# syntax=docker/dockerfile:1
FROM golang:1.17-alpine as dev
WORKDIR /app
RUN apk --no-cache add curl
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s && cp ./bin/air /bin/air
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ENTRYPOINT [ "air" ]

FROM golang:1.17-alpine as builder
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

FROM alpine as prod
WORKDIR /app
COPY --from=builder /build/main ./main
EXPOSE 8000
ENTRYPOINT [ "./main" ]
