package routing

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/mgerb/chi_auth_server/config"
	"github.com/mgerb/chi_auth_server/routing/user"
	"github.com/mgerb/chi_auth_server/util"
)

// Init - initialize routes
func Init() *chi.Mux {

	// new router
	r := chi.NewRouter()

	// default middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultCompress)

	if !config.Flags.Prod {
		r.Use(middleware.Logger)
	}

	// public routes
	r.Group(func(r chi.Router) {
		r.Post("/user/login", user.Login)
		r.Post("/user/create", user.Create)
	})

	// authenticated routes
	r.Group(func(r chi.Router) {

		r.Use(jwtauth.Verifier(util.TokenAuth))

		r.Use(jwtauth.Authenticator)

		// end points
		r.Get("/user/info", user.Info)
	})

	return r
}
