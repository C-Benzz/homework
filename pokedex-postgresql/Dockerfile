# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine


# ENV GOROOT /usr/local/go
# ENV GOCACHE /gocache
# ENV GOPATH /go

# ENV PATH="/go/bin:/usr/local/go/bin:${PATH}"
# Download necessary Go modules
# RUN set -eux; \
# 	export GOROOT="$(go env GOROOT)"; \
# 	./run-tool-which-requires-GOROOT-set.sh
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download github.com/rs/zerolog
# RUN go mod download

COPY . /app
WORKDIR /app
RUN go mod download
# ENV GOROOT /usr/local/go
RUN go build -o ./outserver .


EXPOSE 8080
CMD ["./outserver"]
# CMD ["go","run","server.go"]

# ... the rest of the Dockerfile is ...
# ...   omitted from this example   ...