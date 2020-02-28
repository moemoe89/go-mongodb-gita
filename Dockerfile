FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-mongodb-golang

WORKDIR /go/src/github.com/moemoe89/practicing-mongodb-golang

COPY . /go/src/github.com/moemoe89/practicing-mongodb-golang

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/practicing-mongodb-golang
