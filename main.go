package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dev-parvej/go-api-starter-sql/config"
	migration "github.com/dev-parvej/go-api-starter-sql/db/migration"
	"github.com/dev-parvej/go-api-starter-sql/middleware"
	"github.com/dev-parvej/go-api-starter-sql/routes"
	"github.com/gorilla/mux"
)

func main() {
	// To handle db related command
	if len(os.Args) > 1 {
		migrateAction := os.Args[1]
		if strings.Contains(migrateAction, "db") {
			migration.Migrate(migrateAction)
			return
		}
	}

	fmt.Println("GO API starter with mysql and docker")
	fmt.Printf(":%s", config.Get("APP_PORT"))
	r := mux.NewRouter()
	/**
	* You can always add multiple handler. There is no limitation
	 */
	r.Use(middleware.ApplyJsonHeader)
	routes.RouteHandler(r)
	routes.UserRouteHandler(r)
	routes.AuthRouteHandler(r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Get("APP_PORT")), r))
}
