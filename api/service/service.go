//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package service

import (
	"github.com/moemoe89/go-mongodb-gita/api/api_struct/form"
	"github.com/moemoe89/go-mongodb-gita/api/api_struct/model"
	"github.com/moemoe89/go-mongodb-gita/api/repository/mongo"

	"context"

	"github.com/sirupsen/logrus"
)

// Service represent the services
type Service interface {
	UserCreate(ctx context.Context, user *form.UserForm) (*model.UserModel, int, error)
	UserFind(ctx context.Context) ([]*model.UserModel, int, error)
	UserFindByID(ctx context.Context, id string) (*model.UserModel, int, error)
	UserUpdate(ctx context.Context, user *form.UserForm, id string) (*model.UserModel, int, error)
	UserDelete(ctx context.Context, id string) (int, error)
}

type implService struct {
	log       *logrus.Entry
	mongoRepo mongo.Repository
}

// NewService will create an object that represent the Service interface
func NewService(l *logrus.Entry, mr mongo.Repository) Service {
	return &implService{log: l, mongoRepo: mr}
}
