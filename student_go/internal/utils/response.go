package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	IsValid bool   `json:"isValid"`
	Error   string `json:"error"`
}

const (
	OK    = "ok"
	ERROR = "error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		IsValid: false,
		Error:   err.Error(),
	}
}
