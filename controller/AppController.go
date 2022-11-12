package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	slice "github.com/dev-parvej/js_array_method"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	languages := []string{"Typescript", "GO", "Nodejs", "PHP", "MySql"}
	json.NewEncoder(w).Encode(slice.Map(languages, func(ln string, index int) string {
		return fmt.Sprintf("%d. %s", index+1, ln)
	}))
}
