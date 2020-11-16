package service

import (
	"countries-in-africa/model"
	"countries-in-africa/repository"
	"database/sql"
	"fmt"
	"log"
)

type CountryService interface {
	AddCountries(countries []model.Country) error
	GetCountries() (error, []model.Country)
}

type countryService struct {
	CountryRepository repository.CountryRepository
}

func NewCountryService(db *sql.DB) *countryService {
	return &countryService{
		CountryRepository: repository.NewCountryRepository(db),
	}
}

func (s *countryService) AddCountries(countries []model.Country) error {
	log.Println("Enter: service.addCountries")
	defer log.Println("Exit:  service.addCountries")

	err, countries := s.CountryRepository.AddCountries(countries)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func (s *countryService) GetCountries() (error, []model.Country) {
	log.Println("Enter: service.getCountries")
	defer log.Println("Exit:  service.getCountries")

	err, countries := s.CountryRepository.GetCountries()
	if err != nil {
		log.Println(err)
		return err, []model.Country{}
	}

	var modelCountries []model.Country

	for _, modelCountry := range countries {
		fmt.Println(modelCountry)
		modelCountries = append(modelCountries, model.EntityToModel(modelCountry))
	}

	return err, modelCountries
}
