package util

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
)

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func JsonEncoder(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func JsonDecoder(r *http.Request, target any) {
	json.NewDecoder(r.Body).Decode(target)
}

func ValidateStruct(form interface{}) error {
	err := validator.New().Struct(form)
	if err == nil {
		return nil
	}

	errors := err.(validator.ValidationErrors)

	if errors != nil {
		return errors
	}

	return nil

}

func ParseInt(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}
