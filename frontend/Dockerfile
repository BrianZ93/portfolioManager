FROM node:16.16.0-alpine as frontend

RUN mkdir -p /app

WORKDIR /app

COPY ./frontend/package*.json ./
COPY ./frontend/yarn.lock ./
RUN yarn install

RUN yarn global add @quasar/cli

COPY ./frontend .
