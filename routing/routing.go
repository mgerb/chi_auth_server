package routing

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	myMiddleware "github.com/mgerb/chi_auth_server/middleware"
	"github.com/mgerb/chi_auth_server/routing/user"
)

// Init - initialize routes
func Init() *chi.Mux {

	// new router
	r := chi.NewRouter()

	// default middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.DefaultCompress)

	// public routes
	r.Group(func(r chi.Router) {
		r.Get("/user/login", user.Login)
		r.Get("/user/createUser", user.Create)
	})

	// authenticated routes
	r.Group(func(r chi.Router) {

		// apply auth middleware
		r.Use(myMiddleware.JWTMiddleware)

		// end points
		r.Get("/user/info", user.Info)
	})

	return r
}
