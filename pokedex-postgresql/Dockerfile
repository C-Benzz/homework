# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine

COPY . /app
WORKDIR /app
RUN go mod download

EXPOSE 8080
CMD ["go","run","server.go"]