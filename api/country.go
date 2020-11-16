package api

import (
	"countries-in-africa/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) addCountries(w http.ResponseWriter, r *http.Request) {
	log.Println("Enter: api.addCountries")
	defer log.Println("Exit:  api.addCountries")

	//TODO json validation

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var countries []model.Country

	if err = json.Unmarshal(body, &countries); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = s.CountryService.AddCountries(countries); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (s *Server) getCountries(w http.ResponseWriter, r *http.Request) {
	log.Println("Enter: api.getCountries")
	defer log.Println("Exit:  api.getCountries")

	err, countries := s.CountryService.GetCountries()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(countries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
