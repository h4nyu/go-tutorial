FROM golang:1-alpine

ARG RUNTIME_DEPS="inotify-tools"
WORKDIR /srv
RUN apk add  $RUNTIME_DEPS
COPY . .
