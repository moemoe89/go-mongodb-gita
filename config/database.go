//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitDB will create a variable that represent the redis.Client
func InitDB() (*mongo.Client, error) {
	ctx := context.Background()
	clientOpts := options.Client().ApplyURI(Configuration.MongoDB.Addr)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("Failed to ping connection to mongoDB: %s", err.Error())
	}

	return client, nil
}
