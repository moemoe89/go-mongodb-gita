//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package mongo

import (
	"github.com/moemoe89/go-mongodb-gita/api/api_struct/model"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	USER_COLLECTION = "user"
)

func (m *mongoRepository) UserCreate(ctx context.Context, data *model.UserModel) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return m.queryInsertOne(ctx, USER_COLLECTION, data)
}

func (m *mongoRepository) UserFind(ctx context.Context, filters bson.M, orderBy string, pagination bool, page, perPage int) ([]*model.UserModel, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var users []*model.UserModel
	err := m.queryFind(ctx, USER_COLLECTION, filters, orderBy, pagination, page, perPage, &users)
	return users, err
}

func (m *mongoRepository) UserFindByID(ctx context.Context, filters bson.M, sort string) (*model.UserModel, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var user *model.UserModel
	err := m.queryFindOne(ctx, USER_COLLECTION, filters, sort, &user)
	return user, err
}

func (m *mongoRepository) UserUpdate(ctx context.Context, filters bson.M, data *model.UserModel, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return m.queryUpdateOne(ctx, USER_COLLECTION, filters, data)
}

func (m *mongoRepository) UserDelete(ctx context.Context, filters bson.M) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return m.queryDeleteOne(ctx, USER_COLLECTION, filters)
}
