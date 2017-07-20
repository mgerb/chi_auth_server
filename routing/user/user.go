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

	w.Write([]byte("Login end point."))

	// check database for user info

}

// Create -
func Create(w http.ResponseWriter, r *http.Request) {

	// create new user
	newUser := model.User{
		Email:  "test email",
		Name:   "test name",
		UserID: "1",
	}

	// get new JWT
	newToken, err := util.GetNewJwt(newUser)

	if err != nil {
		response.ERR(w, http.StatusInternalServerError, []byte("Internal error."))
		return
	}

	newUser.Token = newToken

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
