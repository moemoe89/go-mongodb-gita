//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// InitDB will create a variable that represent the redis.Client
func InitDB() (*mgo.Session, error) {
	mongoDBSession, err := mgo.Dial(Configuration.MongoDB.Addr)
	if err != nil {
		return nil, fmt.Errorf("Failed to ping connection to mongoDB: %s", err.Error())
	}

	mongoDBSession.SetSafe(&mgo.Safe{})

	return mongoDBSession, nil
}
