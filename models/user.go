package models

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type User struct {
	Id				bson.ObjectId	`json:"id" bson:"_id"`
	Name			string			`json:"name" bson:"name"`
	Gender			string			`json:"gender" bson:"gender"`
	Age				int				`json:"age" bson:"age"`
	CreatedAt		time.Time		`json:"-" bson:"created_at"`
	UpdatedAt		time.Time		`json:"-" bson:"updated_at"`
}

type UserRequest struct {
	Name	string		`json:"name"`
	Gender	string		`json:"gender"`
	Age		int			`json:"age"`
}

func UserGet()([]User,error){

	dbsession := mongodbSession.Copy()
	defer dbsession.Close()

	c := dbsession.DB(DB_NAME).C(USER_COLLECTION)
	var users []User
	err := c.Find(bson.M{}).All(&users)
	if err != nil {
		return users,err
	}

	return users,nil
}

func UserInsert(user User)(User,error){

	user.Id = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	dbsession := mongodbSession.Copy()
	defer dbsession.Close()

	c := dbsession.DB(DB_NAME).C(USER_COLLECTION)
	err := c.Insert(&user)
	if err != nil {
		return user,err
	}

	return user,nil
}

func UserDetail(id string)(User,error){

	dbsession := mongodbSession.Copy()
	defer dbsession.Close()

	c := dbsession.DB(DB_NAME).C(USER_COLLECTION)
	var user User
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	if err != nil {
		return user,err
	}

	return user,nil
}

func UserUpdate(id string,user User)(User,error){

	userOld,err := UserDetail(id)
	if err != nil {
		return user,err
	}

	user.Id = userOld.Id
	user.CreatedAt = userOld.CreatedAt
	user.UpdatedAt = time.Now()

	dbsession := mongodbSession.Copy()
	defer dbsession.Close()

	c := dbsession.DB(DB_NAME).C(USER_COLLECTION)
	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &user)
	log.Print(err)
	if err != nil {
		return user,err
	}

	return user,nil
}

func UserDelete(id string)(error){

	dbsession := mongodbSession.Copy()
	defer dbsession.Close()

	c := dbsession.DB(DB_NAME).C(USER_COLLECTION)
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return err
	}

	return nil
}