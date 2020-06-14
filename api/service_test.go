//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package api_test

import (
	"github.com/moemoe89/go-mongodb-gita/api/api_struct/form"
	"github.com/moemoe89/go-mongodb-gita/api/api_struct/model"
	ap "github.com/moemoe89/go-mongodb-gita/api"
	"github.com/moemoe89/go-mongodb-gita/api/mocks"
	"github.com/moemoe89/go-mongodb-gita/config"

	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestServiceCreate(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	reqUser := &form.UserForm{
		Name:   "Momo",
		Gender: "male",
		Age:    30,
	}

	mockUser := &model.UserModel{
		ID:     primitive.NewObjectID(),
		Name:   reqUser.Name,
		Gender: reqUser.Gender,
		Age:    reqUser.Age,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*model.UserModel")).Return(mockUser, nil).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.Create(reqUser)

		assert.NoError(t, err)
		assert.NotNil(t, row)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*model.UserModel")).Return(nil, errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.Create(reqUser)

		assert.Error(t, err)
		assert.Nil(t, row)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceFindByID(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)
	mockUser := &model.UserModel{
		ID:        primitive.NewObjectID(),
		Name:      "Momo",
		Gender:    "male",
		Age:       30,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindByID", mock.AnythingOfType("string")).Return(mockUser, nil).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.FindByID(mockUser.ID.Hex())

		assert.NoError(t, err)
		assert.NotNil(t, row)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-not-found", func(t *testing.T) {
		mockRepo.On("FindByID", mock.AnythingOfType("string")).Return(nil, mongo.ErrNoDocuments).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.FindByID(mockUser.ID.Hex())

		assert.Error(t, err)
		assert.Nil(t, row)
		assert.Equal(t, http.StatusNotFound, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("FindByID", mock.AnythingOfType("string")).Return(nil, errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		userRow, status, err := u.FindByID(mockUser.ID.Hex())

		assert.Error(t, err)
		assert.Nil(t, userRow)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceFind(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)
	mockUser := &model.UserModel{
		ID:        primitive.NewObjectID(),
		Name:      "Momo",
		Gender:    "male",
		Age:       30,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	mockListUser := make([]*model.UserModel, 0)
	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Find").Return(mockListUser, nil).Once()
		u := ap.NewService(log, mockRepo)

		users, status, err := u.Find()

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-get", func(t *testing.T) {
		mockRepo.On("Find").Return(nil, errors.New("Unexpected database error")).Once()

		u := ap.NewService(log, mockRepo)

		users, status, err := u.Find()

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceUpdate(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	reqUser := &form.UserForm{
		Name:   "Momo",
		Gender: "male",
		Age:    30,
	}

	mockUser := &model.UserModel{
		ID:     primitive.NewObjectID(),
		Name:   reqUser.Name,
		Gender: reqUser.Gender,
		Age:    reqUser.Age,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindByID", mockUser.ID.Hex()).Return(mockUser, nil).Once()
		mockRepo.On("Update", mockUser, mockUser.ID.Hex()).Return(mockUser, nil).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.Update(reqUser, mockUser.ID.Hex())

		assert.NoError(t, err)
		assert.NotNil(t, row)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("FindByID", mockUser.ID.Hex()).Return(mockUser, nil).Once()
		mockRepo.On("Update", mockUser, mockUser.ID.Hex()).Return(nil, errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.Update(reqUser, mockUser.ID.Hex())

		assert.Error(t, err)
		assert.Nil(t, row)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-detail", func(t *testing.T) {
		mockRepo.On("FindByID", mockUser.ID.Hex()).Return(nil, errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.Update(reqUser, mockUser.ID.Hex())

		assert.Error(t, err)
		assert.Nil(t, row)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceDelete(t *testing.T) {
	id := primitive.NewObjectID()
	log := config.InitLog()
	mockRepo := new(mocks.Repository)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(nil).Once()
		u := ap.NewService(log, mockRepo)

		status, err := u.Delete(id.Hex())

		assert.NoError(t, err)
		assert.Equal(t, 0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		status, err := u.Delete(id.Hex())

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})

}