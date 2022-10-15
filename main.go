package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dev-parvej/go-api-starter/config"
	"github.com/dev-parvej/go-api-starter/controller"
	db "github.com/dev-parvej/go-api-starter/db/migration"
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

	r.HandleFunc("/", controller.ServeHome)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Get("APP_PORT")), r))
}
