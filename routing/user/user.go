package user

import (
	"log"
	"net/http"

	"github.com/mgerb/chi_auth_server/db"
	"github.com/mgerb/chi_auth_server/model"
	"github.com/mgerb/chi_auth_server/response"
	"github.com/mgerb/chi_auth_server/util/jwt"
)

// Login -
func Login(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	response.JSON(w, map[string]interface{}{
		"username": username,
		"password": password,
	})

}

// Create -
func Create(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	// password := r.FormValue("password")

	// TODO - validation

	// TODO - perform database operations
	_, err := db.Conn.Exec(`INSERT INTO users(email, password) VALUES($1, $2);`, email, "password 123")

	if err != nil {
		log.Println(err)
		response.ERR(w, http.StatusInternalServerError, response.DefaultInternalError)
		return
	}

	// create new claims for jwt
	claims := jwt.Claims{
		Email:  email,
		UserID: "1", // temp user id
	}

	// get new JWT
	newToken, exp, err := jwt.GetNewJwt(claims)

	if err != nil {
		response.ERR(w, http.StatusInternalServerError, []byte("Internal error."))
		return
	}

	newUser := model.User{
		Email:  email,
		UserID: "1",
		Token:  newToken,
		Exp:    exp,
	}

	response.JSON(w, newUser)
}

// TokenRefresh - get a new token with fresh expiration date
func TokenRefresh(w http.ResponseWriter, r *http.Request) {

	claims, err := jwt.ParseClaims(r.Context())

	if err != nil {
		response.ERR(w, http.StatusInternalServerError, response.DefaultInternalError)
	}

	newToken, exp, err := jwt.GetNewJwt(claims)

	if err != nil {
		response.ERR(w, http.StatusInternalServerError, response.DefaultInternalError)
	}

	newUser := model.User{
		Email:  claims.Email,
		UserID: claims.UserID,
		Token:  newToken,
		Exp:    exp,
	}

	response.JSON(w, newUser)
}

// Info -
func Info(w http.ResponseWriter, r *http.Request) {

	claims, err := jwt.ParseClaims(r.Context())

	if err != nil {
		response.ERR(w, http.StatusInternalServerError, response.DefaultInternalError)
	}

	response.JSON(w, claims)
}
