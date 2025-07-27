package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gurpreet-fe/go-crud/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	r := chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
	 ______     ______        ______     ______   __
	/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \
	\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \
	 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\
	  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/
	`)
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
