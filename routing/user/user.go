package user

import (
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/mgerb/chi_auth_server/model"
	"github.com/mgerb/chi_auth_server/response"
	"github.com/mgerb/chi_auth_server/util"
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

	// create new user
	newUser := model.User{
		Email:  email,
		UserID: "1", // temp user id
	}

	// get new JWT
	newToken, exp, err := util.GetNewJwt(newUser)

	if err != nil {
		response.ERR(w, http.StatusInternalServerError, []byte("Internal error."))
		return
	}

	newUser.Token = newToken
	newUser.Exp = exp

	response.JSON(w, newUser)
}

// Info -
func Info(w http.ResponseWriter, r *http.Request) {
	// get JWT claims from context that was added in the middleware

	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		response.ERR(w, http.StatusInternalServerError, response.DefaultInternalError)
	}

	email, _ := claims.Get("email")
	name, _ := claims.Get("name")
	userID, _ := claims.Get("userID")

	w.Write([]byte(email.(string) + " " + name.(string) + " " + userID.(string)))
}
