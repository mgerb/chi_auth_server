package user

import (
	"net/http"

	"github.com/mgerb/chi_auth_server/model"
	"github.com/mgerb/chi_auth_server/response"
	"github.com/mgerb/chi_auth_server/util"
)

/*
func Example(w http.ResponseWriter, r *http.Request) {

	// get custom values set by middleware
	s, ok := r.Context().Value("jwt").(*middleware.Claims)

	if !ok {
		log.Println("not ok")
	}

	// get url param set by framework /user/test/{id}
	param := chi.URLParam(r, "test")
}
*/

// Login -
func Login(w http.ResponseWriter, r *http.Request) {

	// check database for user info

	// create new user
	newUser := model.User{
		Email: "test email",
		Name:  "test name",
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

// Create -
func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user end point."))
}

// Info -
func Info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User info end point."))
}
