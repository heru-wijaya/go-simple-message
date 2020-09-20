package controller

import (
	"go-simple-message/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMessageInput is validation for create message body
type CreateMessageInput struct {
	Body string `json:"body" binding:"required"`
}

// GetAllMessage to retrieve all message from db
func GetAllMessage(c *gin.Context) {
	var message []model.Message
	model.DB.Find(&message)

	c.JSON(http.StatusOK, gin.H{"data": message})
}

// CreateMessage to create a new message on database
func CreateMessage(c *gin.Context) {
	var input CreateMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := model.Message{Body: input.Body}
	model.DB.Create(&message)

	c.JSON(http.StatusOK, gin.H{"data": message})
}

// CreateMessageForChat to create a new message on database for chat
func CreateMessageForChat(input CreateMessageInput) {
	message := model.Message{Body: input.Body}
	model.DB.Create(&message)
}
