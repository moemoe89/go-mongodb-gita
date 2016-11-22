package models

import (
	"gopkg.in/mgo.v2"
)

const (
	DB_NAME = "practicing-golang"
	USER_COLLECTION = "user"
)

var mongodbSession *mgo.Session

func init(){
	var err error

	mongodbSession, err = mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	mongodbSession.SetSafe(&mgo.Safe{})
}