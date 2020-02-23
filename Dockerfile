FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-mongodb-golang

WORKDIR /go/src/github.com/moemoe89/practicing-mongodb-golang

COPY . /go/src/github.com/moemoe89/practicing-mongodb-golang

RUN go get gopkg.in/go-playground/validator.v10
RUN mkdir -p /go/src/github.com/moemoe89/practicing-mongodb-golang/vendor/github.com/go-playground/validator/v10
RUN cp -rf /go/src/gopkg.in/go-playground/validator.v10/* /go/src/github.com/moemoe89/practicing-mongodb-golang/vendor/github.com/go-playground/validator/v10
RUN go install

ENTRYPOINT /go/bin/practicing-mongodb-golang
