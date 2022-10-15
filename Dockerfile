FROM golang:latest

WORKDIR /go-api-starter

ADD . .

RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon@latest

ENTRYPOINT CompileDaemon -command="./go-api-starter"