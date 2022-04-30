package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Payload struct {
	Advice string `json:"advice"`
}

func (s *Store) AdviceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rAdvice, err := s.retrieveRandomAdvice()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res := Payload{
			Advice: rAdvice,
		}

		json.NewEncoder(w).Encode(res)
	case http.MethodDelete:
		var rp Payload
		err := json.NewDecoder(r.Body).Decode(&rp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if rp.Advice == "" {
			// check if the value was populated and/or there in the first place
			http.Error(w, "please correct the request body provided", http.StatusBadRequest)
			return
		}

		err = s.remove(rp.Advice)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(map[string]bool{
			"success": true,
		})
	case http.MethodPost:
		var rp Payload
		err := json.NewDecoder(r.Body).Decode(&rp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if rp.Advice == "" {
			// check if the value was populated and/or there in the first place
			http.Error(w, "please correct the request body provided", http.StatusBadRequest)
			return
		}

		err = s.add(rp.Advice)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		json.NewEncoder(w).Encode(map[string]bool{
			"success": true,
		})
	}
}
