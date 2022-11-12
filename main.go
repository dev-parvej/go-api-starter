package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dev-parvej/go-api-starter-sql/config"
	db "github.com/dev-parvej/go-api-starter-sql/db/migration"
	"github.com/dev-parvej/go-api-starter-sql/routes"
	"github.com/gorilla/mux"
)

func main() {
	if len(os.Args) > 1 {
		migrateAction := os.Args[1]
		if strings.Contains(migrateAction, "db") {
			db.Migrate(migrateAction)
			return
		}
	}

	fmt.Println("GO API starter with mysql and docker")
	fmt.Printf(":%s", config.Get("APP_PORT"))
	r := mux.NewRouter()
	/**
	* You can always add multiple handler. There is no limitation
	 */
	routes.RouteHandler(r)
	routes.UserRouteHandler(r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Get("APP_PORT")), r))
}
