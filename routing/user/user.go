package user

import (
	"net/http"

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
	user, ok := r.Context().Value("claims").(*model.User)

	if !ok {
		response.ERR(w, http.StatusInternalServerError, response.DefaultInternalError)
	}

	w.Write([]byte(user.Email + " " + user.Name + " " + user.UserID))
}
