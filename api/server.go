package api

import (
	"countries-in-africa/service"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Router         *mux.Router
	CountryService service.CountryService
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		Router:         mux.NewRouter(),
		CountryService: service.NewCountryService(db),
	}
}

func Run(db *sql.DB) {
	s := NewServer(db)
	s.setHandlers()
	log.Printf("Starting server on %s", ":8090")
	err := http.ListenAndServe(":8090", s.Router)
	log.Fatal(err)
}
