package model

// User -
type User struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Token  string `json:"token"` // JWT token
	UserID string `json:"userID"`
}
