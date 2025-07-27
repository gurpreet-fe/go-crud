package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gurpreet-fe/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	err := http.ListenAndServe("localhost:800", r)
	if err != nil {
		log.Error(err)
	}
}
