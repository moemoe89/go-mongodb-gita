//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Gender    string             `json:"gender" bson:"gender"`
	Age       int                `json:"age" bson:"age"`
	CreatedAt time.Time          `json:"-" bson:"created_at"`
	UpdatedAt time.Time          `json:"-" bson:"updated_at"`
}
