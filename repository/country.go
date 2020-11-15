package repository

import (
	"context"
	"countries-in-africa/model"
	"database/sql"
	"fmt"
	"log"
)

type CountryRepository interface {
	AddCountry(country model.Country) (error, model.Country)
}

type countryRepository struct {
	Database *sql.DB
}

func NewCountryRepository(db *sql.DB) *countryRepository {
	return &countryRepository{
		Database: db,
	}
}

func (s *countryRepository) AddCountry(country model.Country) (error, model.Country) {
	log.Println("Enter: repository.addCountry")
	defer log.Println("Exit:  repository.addCountry")

	sqlInsertCountry := `INSERT INTO country(name, acronym, capital, calling_code, latitude, longitude) values ($1, $2, $3, $4, $5, $6)`
	sqlInsertCurrency := `INSERT INTO currency(country_name, code, name) values ($1,$2,$3)`

	ctx := context.Background()
	tx, err := s.Database.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return err, country
	}

	_, err = tx.Exec(sqlInsertCountry, country.Name, country.Acronym, country.Capital, country.CallingCode, fmt.Sprintf("%f", country.Latitude), fmt.Sprintf("%f", country.Longitude))
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err, country
	}

	for _, currency := range country.Currencies {
		_, err = tx.Exec(sqlInsertCurrency, country.Name, currency.Name, currency.Code)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return err, country
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err, country
	}
	return err, country
}
