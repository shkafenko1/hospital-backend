package models

import "time"

type Patient struct {
	ID          int64     `json:"id"`
	Passport    string    `json:"passport"`
	FirstName   string    `json:"first_name"`
	MiddleName  string    `json:"middle_name"`
	LastName    string    `json:"last_name"`
	Phone       string    `json:"phone"`
	BirthDate   time.Time `json:"birth_date"`
	Sex         string    `json:"sex"`
	Address     string    `json:"address"`
}
