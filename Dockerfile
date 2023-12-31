FROM golang:1.19.0

WORKDIR /usr/src/dockerpractice

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy