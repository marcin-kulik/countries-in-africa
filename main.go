package main

import (
	"countries-in-africa/api"
	"countries-in-africa/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := DBInit()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	api.Run(db)
}
func DBInit() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Port, config.Host, config.User, config.Password, config.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("UNABLE TO OPEN")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("Failed connection to DB")
		panic(err)

	}

	log.Println("Successfully connected to DB")
	return db, nil
}
