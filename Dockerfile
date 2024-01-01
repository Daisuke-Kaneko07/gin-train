FROM golang:latest

WORKDIR /go/src
RUN go mod init gin-train
RUN go mod tidy
RUN go get github.com/gin-gonic/gin
