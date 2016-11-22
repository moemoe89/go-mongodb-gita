package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"practicing-mongodb-golang/models"
)

func UserGet(c *gin.Context){

	user,err := models.UserGet()
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponseData(c,http.StatusOK,"User list.",true,user)
}

func UserAdd(c *gin.Context){

	var user models.User
	var userRequest models.UserRequest
	c.BindJSON(&userRequest)

	user.Name = userRequest.Name
	user.Gender = userRequest.Gender
	user.Age = userRequest.Age

	user,err := models.UserInsert(user)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponseData(c,http.StatusCreated,"Success insert data.",true,user)
}

func UserDetail(c *gin.Context){

	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		JSONResponse(c,http.StatusBadRequest,"Invalid parameter id: not a valid ObjectId",false)
		return
	}

	user,err := models.UserDetail(id)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponseData(c,http.StatusOK,"User detail.",true,user)
}

func UserEdit(c *gin.Context){

	id := c.Param("id")

	var user models.User
	var userRequest models.UserRequest
	c.BindJSON(&userRequest)

	user.Name = userRequest.Name
	user.Gender = userRequest.Gender
	user.Age = userRequest.Age

	user,err := models.UserUpdate(id,user)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponseData(c,http.StatusOK,"Success update data",true,user)
}

func UserDelete(c *gin.Context){

	id := c.Param("id")

	err := models.UserDelete(id)
	if err != nil {
		JSONResponse(c,http.StatusInternalServerError,err.Error(),false)
		return
	}

	JSONResponse(c,http.StatusOK,"Success delete data.",true)
}

func JSONResponse(c *gin.Context,httpStatus int,message string,status bool){
	c.IndentedJSON(httpStatus, gin.H{
		"message": message,
		"status": status,
	})
}

func JSONResponseData(c *gin.Context,httpStatus int,message string,status bool,data interface{}){
	c.IndentedJSON(httpStatus, gin.H{
		"data": data,
		"message": message,
		"status": status,
	})
}