[![Build Status](https://travis-ci.org/moemoe89/practicing-mongodb-golang.svg?branch=master)](https://travis-ci.org/moemoe89/practicing-mongodb-golang)
[![Coverage Status](https://coveralls.io/repos/github/moemoe89/practicing-mongodb-golang/badge.svg?branch=master)](https://coveralls.io/github/moemoe89/practicing-mongodb-golang?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/practicing-mongodb-golang)](https://goreportcard.com/report/github.com/moemoe89/practicing-mongodb-golang)

# PRACTICING-MONGODB-GOLANG #

Practicing MongoDB Using Golang (Gin Gonic Framework) as Programming Language, MongoDB as Database

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ practicing-mongodb-golang/
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

## Setup and Build

* Setup Golang <https://golang.org/>
* Setup MongoDB <https://www.mongodb.com/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> practicing-mongodb-golang
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