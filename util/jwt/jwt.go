package jwt

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/mgerb/chi_auth_server/config"
)

const sixtyDays time.Duration = time.Hour * 24 * 60

// TokenAuth - used to encode/decode tokens
var TokenAuth = jwtauth.New("HS256", []byte(config.Config.JWTSecret), nil)

// Claims - my custom jwt claims
type Claims struct {
	Email  string
	UserID string
	Exp    int64
}

// GetNewJwt - get a new JWT with the expiry time
func GetNewJwt(claims Claims) (string, int64, error) {

	newClaims := jwtauth.Claims{
		"email":  claims.Email,
		"userID": claims.UserID,
	}

	newClaims.SetIssuedNow()

	// sets the exp claims
	newClaims.SetExpiryIn(sixtyDays)

	_, tokenString, err := TokenAuth.Encode(newClaims)

	exp, ok := newClaims.Get("exp")

	if !ok {
		err := errors.New("bad expiration time")
		log.Println(err)
		return "", 0, err
	}

	return tokenString, exp.(int64), err
}

// ParseClaims - parse claims from request context
func ParseClaims(context context.Context) (Claims, error) {

	// get JWT claims from context that was added in the middleware
	_, claims, err := jwtauth.FromContext(context)

	if err != nil {
		return Claims{}, err
	}

	email, ok := claims.Get("email")

	if !ok {
		return Claims{}, err
	}

	userID, ok := claims.Get("userID")

	if !ok {
		return Claims{}, err
	}

	exp, ok := claims.Get("exp")

	if !ok {
		return Claims{}, err
	}

	return Claims{
		Email:  email.(string),
		UserID: userID.(string),
		Exp:    int64(exp.(float64)),
	}, nil
}
