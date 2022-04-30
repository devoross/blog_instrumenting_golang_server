package api

import (
	"encoding/json"
	"net/http"
)

// we're gonna return a random advice slip from : https://api.adviceslip.com/advice
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	advice, err := retrieveAdviceSlip()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(advice.Slip)
}
