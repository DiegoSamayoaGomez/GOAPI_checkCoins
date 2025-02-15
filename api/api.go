package api

import (
	"encoding/json"
	"net/http"
)

// COIN BALANCE PARAMS
type CoinBalanceParams struct {
	Username string
}

// COIN BALANCE RESPONSE (when succesful response)
type CoinBalanceResponse struct {
	Code int

	//ACCOUNT BALANCE
	Balance int64
}

// ERROR RESPONSE
type Error struct {
	//Error code
	Code int
	//Error message
	Message string
}

// Gives an error to the user who sent the request
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
