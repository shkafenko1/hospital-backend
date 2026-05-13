package repository

import (
	"database/sql"
	"hospital-backend/internal/domain/models"
)

type DoctorRepository struct {
	db *sql.DB
}

func NewDoctorRepository(db *sql.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) Create(doctor *models.Doctor) error {
	query := `INSERT INTO doctor (first_name, middle_name, last_name, specialization, office, license_number, phone, experience) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	return r.db.QueryRow(query, doctor.FirstName, doctor.MiddleName, doctor.LastName, 
		doctor.Specialization, doctor.Office, doctor.LicenseNumber, doctor.Phone, doctor.Experience).Scan(&doctor.ID)
}

func (r *DoctorRepository) GetByID(id int64) (*models.Doctor, error) {
	query := `SELECT id, first_name, middle_name, last_name, specialization, office, license_number, phone, experience 
			  FROM doctor WHERE id = $1`
	doctor := &models.Doctor{}
	err := r.db.QueryRow(query, id).Scan(&doctor.ID, &doctor.FirstName, &doctor.MiddleName, 
		&doctor.LastName, &doctor.Specialization, &doctor.Office, &doctor.LicenseNumber, 
		&doctor.Phone, &doctor.Experience)
	if err != nil {
		return nil, err
	}
	return doctor, nil
}

func (r *DoctorRepository) GetAll() ([]models.Doctor, error) {
	query := `SELECT id, first_name, middle_name, last_name, specialization, office, license_number, phone, experience FROM doctor`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []models.Doctor
	for rows.Next() {
		var doctor models.Doctor
		err := rows.Scan(&doctor.ID, &doctor.FirstName, &doctor.MiddleName, &doctor.LastName,
			&doctor.Specialization, &doctor.Office, &doctor.LicenseNumber, &doctor.Phone, &doctor.Experience)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}

func (r *DoctorRepository) Update(doctor *models.Doctor) error {
	query := `UPDATE doctor SET first_name=$1, middle_name=$2, last_name=$3, specialization=$4, 
			  office=$5, license_number=$6, phone=$7, experience=$8 WHERE id=$9`
	_, err := r.db.Exec(query, doctor.FirstName, doctor.MiddleName, doctor.LastName,
		doctor.Specialization, doctor.Office, doctor.LicenseNumber, doctor.Phone, doctor.Experience, doctor.ID)
	return err
}

func (r *DoctorRepository) Delete(id int64) error {
	query := `DELETE FROM doctor WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
