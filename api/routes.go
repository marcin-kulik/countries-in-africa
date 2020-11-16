package api

func (s *Server) setHandlers() {
	s.Router.HandleFunc("/addCountries", s.addCountries).Methods("POST")
	s.Router.HandleFunc("/getCountries", s.getCountries).Methods("GET")
}
