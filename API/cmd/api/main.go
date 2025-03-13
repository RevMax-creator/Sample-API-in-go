package main

import (
	"fmt"
	"net/http"

	"API/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)        // Setting up the logger
	var r *chi.Mux = chi.NewRouter() // It returns a pointer to a mux type which is just a stuct used to setup our api
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
	 ______     _______
	/\  ___\   /\   __ \
	\ \ \__ \  \ \  \/\ \
	 \ \_____\  \ \______\
	  \/_____/   \/______/
	`)
	fmt.Println("Running GO API at localhost:8080")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
