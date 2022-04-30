package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Advice string `json:"advice"`
}

type payload struct {
	Advice        string `json:"advice"`
	UpdatedAdvice string `json:"updated_advice"`
}

func (s *Store) AdviceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rAdvice, err := s.retrieveRandomAdvice()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res := response{
			Advice: rAdvice,
		}

		json.NewEncoder(w).Encode(res)
	case http.MethodDelete:
		var rp payload
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
		var rp payload
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
	case http.MethodPut:
		var rp payload
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

		err = s.update(rp.Advice, rp.UpdatedAdvice)
		if err != nil {
			switch err.(type) {
			case *notFoundError:
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			case *conflictError:
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}

		}

		json.NewEncoder(w).Encode(map[string]bool{
			"success": true,
		})
	}
}
