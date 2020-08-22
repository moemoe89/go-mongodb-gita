//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package http

import (
	"github.com/moemoe89/go-mongodb-gita/api/api_struct/form"
	"github.com/moemoe89/go-mongodb-gita/api/api_struct/model"
	"github.com/moemoe89/go-mongodb-gita/api/service"
	cons "github.com/moemoe89/go-mongodb-gita/constant"

	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ctrl struct represent the delivery for controller
type ctrl struct {
	log *logrus.Entry
	svc service.Service
}

// NewCtrl will create an object that represent the ctrl struct
func NewCtrl(log *logrus.Entry, svc service.Service) *ctrl {
	return &ctrl{log, svc}
}

var starTime = time.Now()

// Ping will handle the ping endpoint
func (ct *ctrl) Ping(c *gin.Context) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	startTime := starTime.In(loc)

	res := map[string]string{
		"start_time": startTime.Format("[02 January 2006] 15:04:05 MST"),
	}
	c.JSON(http.StatusOK, res)
}

// UserCreate will handle the create user endpoint
func (ct *ctrl) UserCreate(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	req := &form.UserForm{}
	if err := c.BindJSON(&req); err != nil {
		ct.log.Errorf("can't get json body: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, []string{"Oops! Something went wrong with your request"}, nil))
		return
	}

	errs := req.Validate()
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, errs, nil))
		return
	}

	user, status, err := ct.svc.UserCreate(ctx, req)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusCreated, model.NewGenericResponse(http.StatusCreated, cons.OK, []string{"Data has been saved"}, user))
}

// UserFind will handle the find user endpoint
func (ct *ctrl) UserFind(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	users, status, err := ct.svc.UserFind(ctx)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been retrieved"}, users))
}

// UserFindByID will handle the find user by id endpoint
func (ct *ctrl) UserFindByID(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id := c.Param("id")

	user, status, err := ct.svc.UserFindByID(ctx, id)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been retrieved"}, user))
}

// UserUpdate will handle the update user endpoint
func (ct *ctrl) UserUpdate(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id := c.Param("id")

	req := &form.UserForm{}
	if err := c.BindJSON(&req); err != nil {
		ct.log.Errorf("can't get json body: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, []string{"Oops! Something went wrong with your request"}, nil))
		return
	}

	errs := req.Validate()
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, errs, nil))
		return
	}

	user, status, err := ct.svc.UserFindByID(ctx, id)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	user, status, err = ct.svc.UserUpdate(ctx, req, id)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been updated"}, user))
}

// UserDelete will handle the delete user endpoint
func (ct *ctrl) UserDelete(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id := c.Param("id")

	status, err := ct.svc.UserDelete(ctx, id)
	if err != nil {
		c.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	c.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been deleted"}, nil))
}
