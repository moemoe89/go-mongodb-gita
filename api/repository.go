//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"github.com/moemoe89/practicing-mongodb-golang/api/api_struct/model"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DB_NAME = "practice_mongodb"
	USER_COLLECTION = "user"
)

// Repository represent the repositories
type Repository interface {
	Create(user *model.UserModel) (*model.UserModel, error)
	Find() ([]*model.UserModel, error)
	FindByID(id string) (*model.UserModel, error)
	Update(user *model.UserModel, id string) (*model.UserModel, error)
	Delete(id string) error
}

type mongoDBRepository struct {
	Client *mongo.Client
}

// NewRedisRepository will create an object that represent the Repository interface
func NewMongoDBRepository(Client *mongo.Client) Repository {
	return &mongoDBRepository{Client}
}

func (m *mongoDBRepository) Create(user *model.UserModel) (*model.UserModel, error) {
	collection := m.Client.Database(DB_NAME).Collection(USER_COLLECTION)
	_, err := collection.InsertOne(context.TODO(), user)
	return user, err
}

func (m *mongoDBRepository) Find() ([]*model.UserModel, error) {
	users := []*model.UserModel{}
	collection := m.Client.Database(DB_NAME).Collection(USER_COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		user := model.UserModel{}
		cursor.Decode(&user)
		users = append(users, &user)
	}
	cursor.Close(context.TODO())
	return users, err
}

func (m *mongoDBRepository) FindByID(id string) (*model.UserModel, error) {
	user := &model.UserModel{}
	objectID, _ := primitive.ObjectIDFromHex(id)
	collection := m.Client.Database(DB_NAME).Collection(USER_COLLECTION)
	err := collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)
	return user, err
}

func (m *mongoDBRepository) Update(user *model.UserModel, id string) (*model.UserModel, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	collection := m.Client.Database(DB_NAME).Collection(USER_COLLECTION)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, bson.M{
		"$set": user,
	})
	return user, err
}

func (m *mongoDBRepository) Delete(id string) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	collection := m.Client.Database(DB_NAME).Collection(USER_COLLECTION)
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}
