package model

// Message model
type Message struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Body string `json:"body"`
}
