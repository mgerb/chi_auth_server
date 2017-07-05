package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mgerb/chi_auth_server/middleware"
)

// Login -
func Login(w http.ResponseWriter, r *http.Request) {
	s, ok := r.Context().Value("jwt").(*middleware.Claims)

	if !ok {
		log.Println("not ok")
	}

	param := chi.URLParam(r, "test")

	fmt.Println(param)

	fmt.Println(s.Email)
	w.Write([]byte("Login end point: " + param))
}

// Create -
func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user end point."))
}

// Info -
func Info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User info end point."))
}
