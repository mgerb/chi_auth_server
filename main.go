package main

import (
	"net/http"

	"github.com/mgerb/chi_auth_server/config"
	"github.com/mgerb/chi_auth_server/routing"
)

func main() {

	// read config file
	config.Init()

	// get new router
	router := routing.Init()

	// start server
	http.ListenAndServe(config.Config.Address, router)
}
