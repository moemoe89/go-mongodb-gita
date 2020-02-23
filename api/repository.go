//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"github.com/moemoe89/practicing-mongodb-golang/api/api_struct/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	MongoDBSession *mgo.Session
}

// NewRedisRepository will create an object that represent the Repository interface
func NewMongoDBRepository(MongoDBSession *mgo.Session) Repository {
	return &mongoDBRepository{MongoDBSession}
}

func (m *mongoDBRepository) Create(user *model.UserModel) (*model.UserModel, error) {
	sess := m.MongoDBSession.Copy()
	defer sess.Close()

	c := sess.DB(DB_NAME).C(USER_COLLECTION)
	err := c.Insert(&user)

	return user, err
}

func (m *mongoDBRepository) Find() ([]*model.UserModel, error) {
	sess := m.MongoDBSession.Copy()
	defer sess.Close()

	c := sess.DB(DB_NAME).C(USER_COLLECTION)
	users := []*model.UserModel{}
	err := c.Find(bson.M{}).All(&users)

	return users, err
}

func (m *mongoDBRepository) FindByID(id string) (*model.UserModel, error) {
	sess := m.MongoDBSession.Copy()
	defer sess.Close()

	user := &model.UserModel{}
	c := sess.DB(DB_NAME).C(USER_COLLECTION)
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)

	return user, err
}

func (m *mongoDBRepository) Update(user *model.UserModel, id string) (*model.UserModel, error) {
	sess := m.MongoDBSession.Copy()
	defer sess.Close()

	c := sess.DB(DB_NAME).C(USER_COLLECTION)
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &user)

	return user, err
}

func (m *mongoDBRepository) Delete(id string) error {
	sess := m.MongoDBSession.Copy()
	defer sess.Close()

	c := sess.DB(DB_NAME).C(USER_COLLECTION)
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
