package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"text-converter/internal/api/controllers/ping"
)

type API struct {
	router *mux.Router
}

func NewAPI() API {
	api := API{
		router: mux.NewRouter(),
	}

	//add routes
	api.router.HandleFunc("/ping", ping.Get).Methods(http.MethodGet)

	//add middlewares
	api.router.Use(loggingMiddleware)
	api.router.Use(authMiddleware)
	return api
}

func (api *API) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

