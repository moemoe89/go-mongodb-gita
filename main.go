package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"practicing-mongodb-golang/controllers"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	r.GET("/ping",controllers.Ping)
	r.GET("/user",controllers.UserGet)
	r.POST("/user",controllers.UserAdd)
	r.GET("/user/:id",controllers.UserDetail)
	r.PUT("/user/:id",controllers.UserEdit)
	r.DELETE("/user/:id",controllers.UserDelete)

	r.Run()
}