package routes

import (
	"github.com/dev-parvej/go-api-starter-sql/controller"
	"github.com/gorilla/mux"
)

func UserRouteHandler(r *mux.Router) {
	userRouter := r.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("/", controller.CreateUser).Methods("POST")
}
