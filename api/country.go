package api

import (
	"countries-in-africa/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) addCountry(w http.ResponseWriter, r *http.Request) {
	log.Println("Enter: api.addCountry")
	defer log.Println("Exit:  api.addCountry")

	//TODO json validation

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	country := model.Country{}

	if err = json.Unmarshal(body, &country); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = s.CountryService.AddCountry(country); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
