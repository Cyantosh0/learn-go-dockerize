FROM golang:alpine

WORKDIR /dockerize-go-app

ADD . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./dockerize-go-app"