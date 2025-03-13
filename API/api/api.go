package api

import (
	"encoding/json"
	"net/http"
)

// Here we write the parameters and responses of our endpoint

// We make stucts representing the parameters and response
type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// Success code, usually 200
	Code int
	// Account Balance
	Balance int64
}

type Error struct {
	// Error code
	Code int
	// Error Message
	Message string
}

// This is the function we use to write an error response to the http response writer
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpexted Error Occured", http.StatusInternalServerError)
	}
)
