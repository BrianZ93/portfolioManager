# Backend

FROM golang:1.18.1-bullseye as builder

RUN mkdir -p /app

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

RUN alias air='$(go env GOPATH)/bin/air'

COPY ./backend/go.mod ./
COPY ./backend/go.sum ./
RUN go mod download

COPY ./backend/Equities.json ./
COPY ./backend/*.go ./