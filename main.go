package main

import (
	"go-simple-message/controller"
	"go-simple-message/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	model.ConnectDataBase()

	r.GET("/message", controller.GetAllMessage)
	r.POST("/message", controller.CreateMessage)

	r.Run()
}
