package models

type Doctor struct {
	ID              int64  `json:"id"`
	FirstName       string `json:"first_name"`
	MiddleName      string `json:"middle_name"`
	LastName        string `json:"last_name"`
	Specialization  string `json:"specialization"`
	Office          int64  `json:"office"`
	LicenseNumber   int64  `json:"license_number"`
	Phone           string `json:"phone"`
	Experience      int64  `json:"experience"`
}
