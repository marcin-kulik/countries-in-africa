package service

import (
	"countries-in-africa/model"
	"countries-in-africa/repository"
	"database/sql"
	"log"
)

type CountryService interface {
	AddCountry(country model.Country) error
}

type countryService struct {
	CountryRepository repository.CountryRepository
}

func NewCountryService(db *sql.DB) *countryService {
	return &countryService{
		CountryRepository: repository.NewCountryRepository(db),
	}
}

func (s *countryService) AddCountry(country model.Country) error {
	log.Println("Enter: service.addCountry")
	defer log.Println("Exit:  service.addCountry")

	err, country := s.CountryRepository.AddCountry(country)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
