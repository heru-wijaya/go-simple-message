package controller

import (
	"go-simple-message/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllMessage to retrieve all message from db
func GetAllMessage(c *gin.Context) {
	var message []model.Message
	model.DB.Find(&message)

	c.JSON(http.StatusOK, gin.H{"data": message})
}
