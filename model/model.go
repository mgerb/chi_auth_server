package model

// User -
type User struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
	Token  string `json:"token"` // JWT token
	Exp    int64  `json:"exp"`   // time when the JWT expires
}
