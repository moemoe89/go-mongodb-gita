//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package routers

import (
	h "github.com/moemoe89/go-mongodb-gita/api/delivery/http"
	mw "github.com/moemoe89/go-mongodb-gita/api/middleware"
	"github.com/moemoe89/go-mongodb-gita/api/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetRouter will create a variable that represent the gin.Engine
func GetRouter(log *logrus.Entry, svc service.Service) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(mw.CORS)

	ctrl := h.NewCtrl(log, svc)
	r.GET("/", ctrl.Ping)
	r.GET("/ping", ctrl.Ping)
	r.POST("/user", ctrl.UserCreate)
	r.GET("/user", ctrl.UserFind)
	r.GET("/user/:id", ctrl.UserFindByID)
	r.PUT("/user/:id", ctrl.UserUpdate)
	r.DELETE("/user/:id", ctrl.UserDelete)

	return r
}
