# build angular
FROM node:16-alpine AS angular_builder
WORKDIR /frontend
COPY frontend/ .
RUN npm install
RUN npm run build


FROM golang:1.18 AS goapi_builder
WORKDIR /backend
COPY ./backend .
RUN go get -d -v
RUN go install -v
RUN go build .


# nginx web server
FROM ubuntu:20.04
RUN apt update && apt install nginx -y
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=angular_builder frontend/dist /usr/share/nginx/html
COPY --from=goapi_builder backend backend/
