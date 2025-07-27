package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/gurpreet-fe/go-crud/api"
	"github.com/gurpreet-fe/go-crud/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params api.CoinBalanceParams
	decoder := schema.NewDecoder()

	err := decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Errorf("Error decoding query parameters: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	database, err := tools.NewDatabase()
	if err != nil {
		log.Errorf("Error creating database connection: %v", err)
		api.InternalErrorHandler(w)
		return
	}

	tokenDetails := (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Errorf("Error getting user coins: %v", err)
		api.InternalErrorHandler(w)
		return
	}

	response := api.CoinBalanceResponse{
		Balance: tokenDetails.Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Errorf("Error encoding response: %v", err)
		api.InternalErrorHandler(w)
		return
	}
}
