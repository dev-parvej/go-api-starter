package routes

import (
	"github.com/dev-parvej/go-api-starter-sql/controller"
	"github.com/gorilla/mux"
)

func RouteHandler(r *mux.Router) {
	r.HandleFunc("/", controller.ServeHome)
}
