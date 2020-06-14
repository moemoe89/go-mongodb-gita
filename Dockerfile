FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-mongodb-gita

WORKDIR /go/src/github.com/moemoe89/go-mongodb-gita

COPY . /go/src/github.com/moemoe89/go-mongodb-gita

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/go-mongodb-gita
