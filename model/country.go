package model

type Country struct {
	Name        string     `json:"name"`
	Acronym     string     `json:"acronym"`
	Capital     string     `json:"capital"`
	CallingCode string     `json:"callingCode"`
	Currencies  []Currency `json:"currencies"`
	Latitude    float32    `json:"latitude"`
	Longitude   float32    `json:"longitude"`
}

type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
