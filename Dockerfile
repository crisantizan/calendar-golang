# Base Stage
FROM golang:1.16.4 as base

# Add Maintainer Info
LABEL maintainer="Cristian Santiz <crisantizan@gmail.com>"

ARG APP_PATH=/app

WORKDIR $APP_PATH

# Copy go mod
COPY go.mod ./

# Download all dependencies.
RUN go mod download && mkdir -p bin

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Development Stage
FROM base as dev

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o bin/calendar-golang -v main.go" --command=./bin/calendar-golang

