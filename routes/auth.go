package routes

import (
	"github.com/dev-parvej/go-api-starter-sql/controller"
	"github.com/gorilla/mux"
)

func AuthRouteHandler(r *mux.Router) {
	authRouter := r.PathPrefix("/api").Subrouter()

	authRouter.HandleFunc("/login", controller.Login).Methods("POST")
}
