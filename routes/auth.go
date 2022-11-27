package routes

import (
	"net/http"

	"github.com/dev-parvej/go-api-starter-sql/controller"
	"github.com/dev-parvej/go-api-starter-sql/middleware"
	"github.com/gorilla/mux"
)

func AuthRouteHandler(r *mux.Router) {
	r.HandleFunc("/api/login", controller.Login).Methods("POST")

	authRouter := r.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.Authenticate)

	authRouter.HandleFunc("/log-out", func(w http.ResponseWriter, r *http.Request) {
		//To Implement log out code here
	})
}
