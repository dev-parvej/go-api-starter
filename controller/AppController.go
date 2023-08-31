package controller

import (
	"fmt"
	"net/http"

	"github.com/dev-parvej/go-api-starter-sql/util"
	slice "github.com/dev-parvej/js_array_method"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	languages := []string{"Typescript", "GO", "Nodejs", "PHP", "MySql", "Dart"}

	util.JsonEncoder(w, slice.Map(languages, func(ln string, index int) string {
		return fmt.Sprintf("%d. %s", index+1, ln)
	}))
}
