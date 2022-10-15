package controller

import (
	"encoding/json"
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to go API starter kit")
}
