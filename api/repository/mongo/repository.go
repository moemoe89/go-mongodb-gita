//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package mongo

import (
	"github.com/moemoe89/go-mongodb-gita/api/api_struct/model"

	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository represent the repositories
type Repository interface {
	UserCreate(ctx context.Context, data *model.UserModel) error
	UserFind(ctx context.Context, filters bson.M, orderBy string, pagination bool, page, perPage int) ([]*model.UserModel, error)
	UserFindByID(ctx context.Context, filters bson.M, sort string) (*model.UserModel, error)
	UserUpdate(ctx context.Context, filters bson.M, data *model.UserModel, id string) error
	UserDelete(ctx context.Context, filters bson.M) error
}

type mongoRepository struct {
	client *mongo.Client
	dbName string
}

// NewMongoRepository will create an object that represent the Repository interface
func NewMongoRepository(client *mongo.Client, dbName string) Repository {
	return &mongoRepository{client, dbName}
}

func (m *mongoRepository) querySort(sort string) bson.M {
	sortDocument := bson.M{}

	split := strings.Split(sort, "")
	if len(split) > 0 {
		if split[0] == "-" {
			sort = strings.Replace(sort, "-", "", -1)
			sortDocument = bson.M{
				sort: -1,
			}
		} else {
			sortDocument = bson.M{
				sort: 1,
			}
		}
	}
	return sortDocument
}

func (m *mongoRepository) queryInsertOne(ctx context.Context, collection string, data interface{}) error {
	c := m.client.Database(m.dbName).Collection(collection)
	_, err := c.InsertOne(ctx, data)
	return err
}

func (m *mongoRepository) queryFind(ctx context.Context, collection string, filters bson.M, orderBy string, pagination bool, page, perPage int, data interface{}) error {
	opts := options.Find()
	if len(orderBy) > 0 {
		opts.SetSort(m.querySort(orderBy))
	}

	if pagination {
		opts.SetSkip(int64(page))
	}

	if perPage > 0 {
		opts.SetLimit(int64(perPage))
	}

	c := m.client.Database(m.dbName).Collection(collection)
	cursor, err := c.Find(ctx, filters, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, data)

	return err
}

func (m *mongoRepository) queryFindOne(ctx context.Context, collection string, filters bson.M, sort string, data interface{}) error {
	opts := options.FindOne()
	if len(sort) > 0 {
		opts.SetSort(m.querySort(sort))
	}

	c := m.client.Database(m.dbName).Collection(collection)
	err := c.FindOne(ctx, filters, opts).Decode(data)
	return err
}

func (m *mongoRepository) queryUpdateOne(ctx context.Context, collection string, filters bson.M, data interface{}) error {
	c := m.client.Database(m.dbName).Collection(collection)
	_, err := c.UpdateOne(ctx, filters, bson.M{
		"$set": data,
	})
	return err
}

func (m *mongoRepository) queryDeleteOne(ctx context.Context, collection string, filters bson.M) error {
	c := m.client.Database(m.dbName).Collection(collection)
	_, err := c.DeleteOne(ctx, filters)
	return err
}
