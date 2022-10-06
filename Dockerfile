############################################################################################
# Build executable binary
############################################################################################
FROM golang:1.18.3-alpine3.15

WORKDIR /app

# copy and fetch external module
COPY go.mod .
RUN go mod download && \
    go mod verify

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o couchbase-test

ENTRYPOINT ["./couchbase-test"]



