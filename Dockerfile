FROM golang:latest

WORKDIR /go-api-starter-sql

ADD . .

RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon@latest

ENTRYPOINT CompileDaemon -command="./go-api-starter-sql"