package util

import (
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/mgerb/chi_auth_server/config"
	"github.com/mgerb/chi_auth_server/model"
)

// TokenAuth - used to encode/decode tokens
var TokenAuth = jwtauth.New("HS256", []byte(config.Config.JWTSecret), nil)

// GetNewJwt - get a new token
func GetNewJwt(user model.User) (string, error) {

	newClaims := jwtauth.Claims{
		"email":  user.Email,
		"name":   user.Name,
		"userID": user.UserID,
	}

	newClaims.SetIssuedNow().SetExpiryIn(time.Hour * 24 * 60)

	_, tokenString, err := TokenAuth.Encode(newClaims)

	return tokenString, err
}
