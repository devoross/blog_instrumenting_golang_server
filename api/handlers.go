package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Payload struct {
	Advice string `json:"advice"`
}

// we're gonna return a random advice slip from : https://api.adviceslip.com/advice
func (s *Store) AdviceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rAdvice := s.retrieveRandomAdvice()
		res := Payload{
			Advice: rAdvice,
		}

		json.NewEncoder(w).Encode(res)
	case http.MethodDelete:
		/*
			1: Parse the incoming payload
			2: Validate payload and return 400 if it fails, continue if OK
			3: Use the incoming payload to delete the resource
			4: If resource not found 404
			5: Else 200
		*/
		var rp Payload
		err := json.NewDecoder(r.Body).Decode(&rp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		err = s.remove(rp.Advice)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	case http.MethodPost:
		/*
			1: Parse the incoming payload
			2: Validate payload and return 400 if it fails, continue if OK
			3: Use the incoming payload to try and create the item
			4: If there's an error, the item already exists in the slice, and therefore what should we return (409 conflict?)
			5: Else 200
		*/
		var rp Payload
		err := json.NewDecoder(r.Body).Decode(&rp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		err = s.add(rp.Advice)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusConflict)
		}
	}
}
