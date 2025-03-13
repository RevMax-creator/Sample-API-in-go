package middleware

import (
	"errors"
	"net/http"

	"API/api"
	"API/internal/tools"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token") // Custom message

// Since this function is acting as an middleware it needs to take in and return a http handler interface
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r) // This calls next middleware in line or the handler function as there's no middlewaare left
		// Authorization -> next.ServeHTTP -> GetCoinBalance
	})
}
