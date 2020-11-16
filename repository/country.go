package repository

import (
	"context"
	"countries-in-africa/entity"
	"countries-in-africa/model"
	"database/sql"
	"fmt"
	"log"
)

type CountryRepository interface {
	AddCountries(countries []model.Country) (error, []model.Country)
	GetCountries() (error, []entity.Country)
}

type countryRepository struct {
	Database *sql.DB
}

func NewCountryRepository(db *sql.DB) *countryRepository {
	return &countryRepository{
		Database: db,
	}
}

func (s *countryRepository) AddCountries(countries []model.Country) (error, []model.Country) {
	log.Println("Enter: repository.AddCountries")
	defer log.Println("Exit:  repository.AddCountries")

	sqlInsertCountry := `INSERT INTO country(name, acronym, capital, calling_code, latitude, longitude) values ($1, $2, $3, $4, $5, $6)`
	sqlInsertCurrency := `INSERT INTO currency(name, code, country_name) values ($1,$2,$3)`

	ctx := context.Background()
	tx, err := s.Database.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return err, countries
	}

	for _, country := range countries {
		_, err = tx.Exec(sqlInsertCountry, country.Name, country.Acronym, country.Capital, country.CallingCode, fmt.Sprintf("%f", country.Latitude), fmt.Sprintf("%f", country.Longitude))
		if err != nil {
			log.Println(err)
			_ = tx.Rollback()
			return err, countries
		}
		for _, currency := range country.Currencies {
			_, err = tx.Exec(sqlInsertCurrency, currency.Name, currency.Code, country.Name)
			if err != nil {
				log.Println(err)
				_ = tx.Rollback()
				return err, countries
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err, countries
	}
	return err, countries
}

func (s *countryRepository) GetCountries() (error, []entity.Country) {
	log.Println("Enter: repository.GetCountries")
	defer log.Println("Exit:  repository.GetCountries")

	var countries []entity.Country

	rows, err := s.Database.Query("SELECT * FROM country")
	if err != nil {
		log.Println(err)
		return err, countries
	}
	defer rows.Close()

	for rows.Next() {
		var country entity.Country
		err = rows.Scan(&country.Name, &country.Acronym, &country.Capital, &country.CallingCode, &country.Latitude, &country.Longitude)
		if err != nil {
			log.Println(err)
			return err, countries
		}
		countries = append(countries, country)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return err, countries
	}

	return err, countries
}
