//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Gender    string        `json:"gender" bson:"gender"`
	Age       int           `json:"age" bson:"age"`
	CreatedAt time.Time     `json:"-" bson:"created_at"`
	UpdatedAt time.Time     `json:"-" bson:"updated_at"`
}