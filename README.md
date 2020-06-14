[![Build Status](https://travis-ci.org/moemoe89/go-mongodb-gita.svg?branch=master)](https://travis-ci.org/moemoe89/go-mongodb-gita)
[![Coverage Status](https://coveralls.io/repos/github/moemoe89/go-mongodb-gita/badge.svg?branch=master)](https://coveralls.io/github/moemoe89/go-mongodb-gita?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/go-mongodb-gita)](https://goreportcard.com/report/github.com/moemoe89/go-mongodb-gita)

# GO-MONGODB-GITA #

Practicing MongoDB Using Golang (Gin Gonic Framework) with Go Mod as Programming Language, MongoDB as Database

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ go-mongodb-gita/
  |     |
  |     +--+ main.go
  |        + api/
  |        + routers/
  |        + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Requirements

Go >= 1.11

## Setup and Build

* Setup Golang <https://golang.org/>
* Setup MongoDB <https://www.mongodb.com/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> go-mongodb-gita
```

## Running Application
Make config file for local :
```
$ cp config-sample.json config-local.json
```
Build
```
$ go build
```
Run
```
$ go run main.go
```

## How to Run with Docker
Make config file for docker :
```
$ cp config-sample.json config-docker.json
```
Build
```
$ docker-compose build
```
Run
```
$ docker-compose up
```
Stop
```
$ docker-compose down
```

## How to Run Unit Test
Run
```
$ go test ./...
```
Run with cover
```
$ go test ./... -cover
```
Run with HTML output
```
$ go test ./... -coverprofile=c.out && go tool cover -html=c.out
```

## License

MIT