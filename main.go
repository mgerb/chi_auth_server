package main

import (
	"net/http"

	"log"

	"github.com/mgerb/chi_auth_server/config"
	"github.com/mgerb/chi_auth_server/routing"
	"golang.org/x/crypto/acme/autocert"
)

func main() {

	// read config file
	config.Init()

	// get new router
	router := routing.Init()

	// start server

	if config.Flags.TLS {

		// start server on port 80 to redirect
		go http.ListenAndServe(":80", http.HandlerFunc(redirect))

		// start TLS server
		log.Fatal(http.Serve(autocert.NewListener(), router))

	} else {

		// start basic server
		http.ListenAndServe(config.Config.Address, router)
	}
}

// redirect to https
func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}

	http.Redirect(w, req, target,
		// see @andreiavrammsd comment: often 307 > 301
		http.StatusTemporaryRedirect)
}
