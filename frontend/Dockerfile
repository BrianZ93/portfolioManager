# Current node version VueJS was built in
FROM node:16.16.0-alpine as frontend

RUN mkdir -p /app

WORKDIR /app

COPY ./blackswan/package*.json ./
COPY ./blackswan/yarn.lock ./
RUN yarn install

RUN yarn global add @quasar/cli

COPY ./blackswan .

EXPOSE $PORT
