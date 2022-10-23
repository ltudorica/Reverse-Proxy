# syntax=docker/dockerfile:1

FROM golang:1.19.2-alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /reverse-proxy-app

EXPOSE 8080

CMD [ "/docker-gs-ping" ]