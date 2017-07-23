package util

import (
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/mgerb/chi_auth_server/config"
	"github.com/mgerb/chi_auth_server/model"
)

const sixtyDays time.Duration = time.Hour * 24 * 60

// TokenAuth - used to encode/decode tokens
var TokenAuth = jwtauth.New("HS256", []byte(config.Config.JWTSecret), nil)

// GetNewJwt - get a new JWT with the expiry time
func GetNewJwt(user model.User) (string, int64, error) {

	newClaims := jwtauth.Claims{
		"email":  user.Email,
		"userID": user.UserID,
	}

	newClaims.SetIssuedNow().SetExpiryIn(sixtyDays)

	_, tokenString, err := TokenAuth.Encode(newClaims)

	exp := time.Now().Unix() + int64(sixtyDays.Seconds())

	return tokenString, exp, err
}
