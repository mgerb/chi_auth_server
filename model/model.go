package model

// User -
type User struct {
	Email string `json:"email"`
	Name  string `json:"Name"`
	Token string `json:"token"` // JWT token
}
