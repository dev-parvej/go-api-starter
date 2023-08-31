package routes

import (
	"github.com/dev-parvej/go-api-starter-sql/controller"
	"github.com/dev-parvej/go-api-starter-sql/middleware"
	"github.com/gorilla/mux"
)

func AuthRouteHandler(r *mux.Router) {
	r.HandleFunc("/api/login", controller.Login).Methods("POST")
	r.HandleFunc("/api/refresh", controller.GrantAccessToken).Methods("post")

	authRouter := r.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.Authenticate)

	authRouter.HandleFunc("/log-out", controller.LogOut).Methods("post")

}
