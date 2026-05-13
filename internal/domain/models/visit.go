package models

import "time"

type Visit struct {
	ID         int64     `json:"id"`
	DateTime   time.Time `json:"date_time"`
	Symptoms   string    `json:"symptoms"`
	Diagnosis  string    `json:"diagnosis"`
	Advise     string    `json:"advise"`
	Price      float64   `json:"price"`
	PatientID  int64     `json:"patient_id"`
	DoctorID   int64     `json:"doctor_id"`
}
