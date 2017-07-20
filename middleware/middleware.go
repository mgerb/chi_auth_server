package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mgerb/chi_auth_server/config"
	"github.com/mgerb/chi_auth_server/model"
	"github.com/mgerb/chi_auth_server/response"
)

// JWTMiddleware -
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := strings.Split(r.Header.Get("Authorization"), " ")

		// check to see if auth header has a Bearer and a Token
		if len(authHeader) != 2 || authHeader[0] != "Bearer" {
			response.ERR(w, http.StatusUnauthorized, response.DefaultUnauthorized)
			return
		}

		tokenString := authHeader[1]

		// parse the incoming token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(config.Config.JWTSecret), nil
		})

		if err != nil {
			log.Println(err)
			response.ERR(w, http.StatusInternalServerError, response.DefaultUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			log.Println(err)
			response.ERR(w, http.StatusInternalServerError, response.DefaultUnauthorized)
			return
		}

		// check all of our claim types first
		if _, ok = claims["email"].(string); !ok {
			log.Println("Email claims error.")
			response.ERR(w, http.StatusInternalServerError, response.DefaultUnauthorized)
			return
		}

		if _, ok = claims["name"].(string); !ok {
			log.Println("Name claims error.")
			response.ERR(w, http.StatusInternalServerError, response.DefaultUnauthorized)
			return
		}

		if _, ok = claims["userID"].(string); !ok {
			log.Println("UserID claims error.")
			response.ERR(w, http.StatusInternalServerError, response.DefaultUnauthorized)
			return
		}

		// create new claims object
		newClaims := &model.User{
			UserID: claims["userID"].(string),
			Email:  claims["email"].(string),
			Name:   claims["name"].(string),
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "claims", newClaims)))
	})
}
