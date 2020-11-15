package api

func (s *Server) setHandlers() {
	s.Router.HandleFunc("/addCountry", s.addCountry).Methods("POST")
	//s.Router.HandleFunc("/getCountries", s.getCountries).Methods("GET")
}
