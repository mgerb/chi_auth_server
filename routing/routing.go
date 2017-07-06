package routing

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	myMiddleware "github.com/mgerb/chi_auth_server/middleware"
	"github.com/mgerb/chi_auth_server/routing/user"
)

// Init - initialize routes
func Init() http.Handler {

	// new router
	r := chi.NewRouter()

	// default middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.DefaultCompress)

	// apply auth middleware
	r.Use(myMiddleware.JWTMiddleware)

	r.Get("/user/login", user.Login)
	r.Get("/user/createUser", user.Create)

	// end points
	r.Get("/user/info", user.Info)

	return r
}
