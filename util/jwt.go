package util

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mgerb/chi_auth_server/config"
	"github.com/mgerb/chi_auth_server/model"
)

// GetNewJwt - get a new token
func GetNewJwt(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  user.Email,
		"name":   user.Name,
		"userID": user.UserID,
		"iat":    time.Now().Unix(), // issued at
	})

	tokenString, err := token.SignedString([]byte(config.Config.JWTSecret))

	log.Println(tokenString)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}
