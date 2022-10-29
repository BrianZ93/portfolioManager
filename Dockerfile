FROM golang:1.18.5 as backend

RUN mkdir /backend
WORKDIR /backend

COPY ./backend ./

RUN export GO111MODULE=on
RUN go get "github.com/sirupsen/logrus"
RUN go get "github.com/piquette/finance-go"

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8081

FROM node:16.16-alpine as frontend

RUN yarn global add @quasar/cli

WORKDIR /frontend
COPY ./frontend/package.json ./
COPY ./frontend/yarn.lock ./
RUN yarn

COPY ./frontend ./

EXPOSE 9000


