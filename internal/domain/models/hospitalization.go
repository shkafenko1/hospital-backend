package models

import "time"

type Hospitalization struct {
	ID                int64     `json:"id"`
	Type              string    `json:"type"`
	Status            string    `json:"status"`
	BeginningDate     time.Time `json:"beginning_date"`
	PrimaryDiagnosis  string    `json:"primary_diagnosis"`
	DischargeDate     *time.Time `json:"discharge_date,omitempty"`
	DischargeDiagnosis *string  `json:"discharge_diagnosis,omitempty"`
	PatientID         int64     `json:"patient_id"`
	DoctorID          int64     `json:"doctor_id"`
	WardID            int64     `json:"ward_id"`
}
