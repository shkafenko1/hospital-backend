package repository

import (
	"database/sql"
	"hospital-backend/internal/domain/models"
)

type PatientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) Create(patient *models.Patient) error {
	query := `INSERT INTO patient (passport, first_name, middle_name, last_name, phone, birth_date, sex, address) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	return r.db.QueryRow(query, patient.Passport, patient.FirstName, patient.MiddleName, 
		patient.LastName, patient.Phone, patient.BirthDate, patient.Sex, patient.Address).Scan(&patient.ID)
}

func (r *PatientRepository) GetByID(id int64) (*models.Patient, error) {
	query := `SELECT id, passport, first_name, middle_name, last_name, phone, birth_date, sex, address 
			  FROM patient WHERE id = $1`
	patient := &models.Patient{}
	err := r.db.QueryRow(query, id).Scan(&patient.ID, &patient.Passport, &patient.FirstName, 
		&patient.MiddleName, &patient.LastName, &patient.Phone, &patient.BirthDate, 
		&patient.Sex, &patient.Address)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func (r *PatientRepository) GetAll() ([]models.Patient, error) {
	query := `SELECT id, passport, first_name, middle_name, last_name, phone, birth_date, sex, address FROM patient`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		err := rows.Scan(&patient.ID, &patient.Passport, &patient.FirstName, &patient.MiddleName,
			&patient.LastName, &patient.Phone, &patient.BirthDate, &patient.Sex, &patient.Address)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	return patients, nil
}

func (r *PatientRepository) Update(patient *models.Patient) error {
	query := `UPDATE patient SET passport=$1, first_name=$2, middle_name=$3, last_name=$4, 
			  phone=$5, birth_date=$6, sex=$7, address=$8 WHERE id=$9`
	_, err := r.db.Exec(query, patient.Passport, patient.FirstName, patient.MiddleName,
		patient.LastName, patient.Phone, patient.BirthDate, patient.Sex, patient.Address, patient.ID)
	return err
}

func (r *PatientRepository) Delete(id int64) error {
	query := `DELETE FROM patient WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
