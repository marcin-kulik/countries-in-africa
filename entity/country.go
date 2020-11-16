package entity

type Country struct {
	Name        string     `db:"name"`
	Acronym     string     `db:"acronym"`
	Capital     string     `db:"capital"`
	CallingCode string     `db:"calling_code"`
	Currencies  []Currency `db:"currencies"`
	Latitude    float32    `db:"latitude"`
	Longitude   float32    `db:"longitude"`
}

type Currency struct {
	Code string `db:"code"`
	Name string `db:"name"`
}
