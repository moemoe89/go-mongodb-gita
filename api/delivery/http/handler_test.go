//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package http_test

import (
	"github.com/moemoe89/practicing-mongodb-golang/api/api_struct/form"
	"github.com/moemoe89/practicing-mongodb-golang/api/api_struct/model"
	"github.com/moemoe89/practicing-mongodb-golang/api/mocks"
	"github.com/moemoe89/practicing-mongodb-golang/config"
	"github.com/moemoe89/practicing-mongodb-golang/routers"

	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeliveryCreate(t *testing.T) {
	log := config.InitLog()

	userForm := &form.UserForm{
		Name:   "Momo",
		Gender: "male",
		Age:    30,
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("Create", userForm).Return(nil, 0, nil)

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryCreateFail(t *testing.T) {
	log := config.InitLog()

	userForm := &form.UserForm{
		Name:   "Momo",
		Gender: "male",
		Age:    30,
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("Create", userForm).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryCreateFailValidation(t *testing.T) {
	log := config.InitLog()

	userForm := &form.UserForm{
		Name:   "",
		Gender: "male",
		Age:    30,
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryCreateFailBindJSON(t *testing.T) {
	log := config.InitLog()

	mockService := new(mocks.Service)

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/user", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryUpdate(t *testing.T) {
	id := primitive.NewObjectID()

	log := config.InitLog()

	userForm := &form.UserForm{
		Name:   "Momo",
		Gender: "male",
		Age:    30,
	}
	user := &model.UserModel{
		ID:     id,
		Name:   userForm.Name,
		Gender: userForm.Gender,
		Age:    userForm.Age,
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("FindByID", id.Hex()).Return(user, 0, nil)
	mockService.On("Update", userForm, id.Hex()).Return(user, 0, nil)

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/user/"+id.Hex(), strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryUpdateFail(t *testing.T) {
	id := primitive.NewObjectID()

	log := config.InitLog()

	userForm := &form.UserForm{
		Name:   "Momo",
		Gender: "male",
		Age:    30,
	}
	user := &model.UserModel{
		ID:     id,
		Name:   userForm.Name,
		Gender: "male",
		Age:    30,
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("FindByID", id.String()).Return(user, 0, nil)
	mockService.On("Update", userForm, id.String()).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/user/"+id.String(), strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryUpdateFailFindByID(t *testing.T) {
	id := primitive.NewObjectID().Hex()

	log := config.InitLog()

	userForm := &form.UserForm{
		Name:   "Momo",
		Gender: "male",
		Age:    30,
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("FindByID", id).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/user/"+id, strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryUpdateFailValidation(t *testing.T) {
	id := primitive.NewObjectID().Hex()

	log := config.InitLog()

	userForm := &form.UserForm{
		Name:   "",
		Gender: "male",
		Age:    30,
	}

	j, err := json.Marshal(userForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/user/"+id, strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryUpdateFailBindJSON(t *testing.T) {
	id := primitive.NewObjectID().Hex()

	log := config.InitLog()

	mockService := new(mocks.Service)

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/user/"+id, strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryFind(t *testing.T) {
	log := config.InitLog()

	user := &model.UserModel{
		ID:     primitive.NewObjectID(),
		Name:   "Momo",
		Gender: "male",
		Age:    30,
	}
	users := []*model.UserModel{}
	users = append(users, user)

	mockService := new(mocks.Service)
	mockService.On("Find").Return(users, 0, nil)

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryFindFail(t *testing.T) {
	log := config.InitLog()

	mockService := new(mocks.Service)
	mockService.On("Find").Return(nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later"))

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryFindByID(t *testing.T) {
	id := primitive.NewObjectID()

	log := config.InitLog()

	user := &model.UserModel{
		ID:     id,
		Name:   "Momo",
		Gender: "male",
		Age:    30,
	}

	mockService := new(mocks.Service)
	mockService.On("FindByID", id.String()).Return(user, 0, nil)

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/"+id.String(), strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestDeliveryFindByIDFail(t *testing.T) {
	id := primitive.NewObjectID().Hex()

	log := config.InitLog()
	mockService := new(mocks.Service)
	mockService.On("FindByID", id).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeliveryDelete(t *testing.T) {
	id := primitive.NewObjectID().Hex()

	log := config.InitLog()
	mockService := new(mocks.Service)
	mockService.On("Delete", id).Return(0, nil)

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeliveryDeleteFail(t *testing.T) {
	id := primitive.NewObjectID().Hex()

	log := config.InitLog()
	mockService := new(mocks.Service)
	mockService.On("Delete", id).Return(http.StatusInternalServerError, errors.New("Unexpected database error"))

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/user/"+id, strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
