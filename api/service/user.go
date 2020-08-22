//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package service

import (
	"github.com/moemoe89/go-mongodb-gita/api/api_struct/form"
	"github.com/moemoe89/go-mongodb-gita/api/api_struct/model"

	"context"
	"errors"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *implService) UserCreate(ctx context.Context, req *form.UserForm) (*model.UserModel, int, error) {
	user := &model.UserModel{
		ID:        primitive.NewObjectID(),
		Name:      req.Name,
		Gender:    req.Gender,
		Age:       req.Age,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err := s.mongoRepo.UserCreate(ctx, user)
	if err != nil {
		s.log.Errorf("can't create data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, http.StatusCreated, nil
}

func (s *implService) UserFind(ctx context.Context) ([]*model.UserModel, int, error) {
	users, err := s.mongoRepo.UserFind(ctx, nil, "", false, 0, 0)
	if err != nil {
		s.log.Errorf("can't get data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return users, http.StatusOK, nil
}

func (s *implService) UserFindByID(ctx context.Context, id string) (*model.UserModel, int, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("Invalid object id")
	}

	user, err := s.mongoRepo.UserFindByID(ctx, bson.M{
		"_id": objectID,
	}, "")
	if err == mongo.ErrNoDocuments {
		return nil, http.StatusNotFound, errors.New("Data not found")
	}

	if err != nil {
		s.log.Errorf("can't get detail data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, http.StatusOK, nil
}

func (s *implService) UserUpdate(ctx context.Context, req *form.UserForm, id string) (*model.UserModel, int, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("Invalid object id")
	}

	user, err := s.mongoRepo.UserFindByID(ctx, bson.M{
		"_id": objectID,
	}, "")
	if err != nil {
		s.log.Errorf("can't get detail data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	user.Name = req.Name
	user.Gender = req.Gender
	user.Age = req.Age
	user.UpdatedAt = time.Now().UTC()

	err = s.mongoRepo.UserUpdate(ctx, bson.M{
		"_id": objectID,
	}, user, "")
	if err != nil {
		s.log.Errorf("can't update data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, http.StatusOK, nil
}

func (s *implService) UserDelete(ctx context.Context, id string) (int, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return http.StatusBadRequest, errors.New("Invalid object id")
	}

	err = s.mongoRepo.UserDelete(ctx, bson.M{
		"_id": objectID,
	})
	if err != nil {
		s.log.Errorf("can't delete data: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return http.StatusOK, nil
}
