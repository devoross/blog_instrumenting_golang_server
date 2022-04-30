package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type slip struct {
	Advice string `json:"advice"`
}

type adviceResponse struct {
	Slip slip `json:"slip"`
}

func retrieveAdviceSlip() (adviceResponse, error) {
	var r adviceResponse

	hClient := http.Client{Timeout: time.Duration(5 * time.Second)}

	request, err := http.NewRequest("GET", "https://api.adviceslip.com/advice", nil)
	if err != nil {
		log.Println(err)
		return adviceResponse{}, err
	}

	resp, err := hClient.Do(request)
	if err != nil {
		log.Println(err)
		return adviceResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return adviceResponse{}, errors.New("http status returned was not ok")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return adviceResponse{}, err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &r)
	if err != nil {
		return adviceResponse{}, err
	}

	return r, nil
}
