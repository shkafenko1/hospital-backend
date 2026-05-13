package repository

import (
	"database/sql"
	"hospital-backend/internal/domain/models"
)

type HospitalizationRepository struct {
	db *sql.DB
}

func NewHospitalizationRepository(db *sql.DB) *HospitalizationRepository {
	return &HospitalizationRepository{db: db}
}

func (r *HospitalizationRepository) Create(hospitalization *models.Hospitalization) error {
	query := `INSERT INTO hospitalization (type, status, beginning_date, primary_diagnosis, discharge_date, discharge_diagnosis, patient_id, doctor_id, ward_id) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	return r.db.QueryRow(query, hospitalization.Type, hospitalization.Status, hospitalization.BeginningDate, 
		hospitalization.PrimaryDiagnosis, hospitalization.DischargeDate, hospitalization.DischargeDiagnosis,
		hospitalization.PatientID, hospitalization.DoctorID, hospitalization.WardID).Scan(&hospitalization.ID)
}

func (r *HospitalizationRepository) GetByID(id int64) (*models.Hospitalization, error) {
	query := `SELECT id, type, status, beginning_date, primary_diagnosis, discharge_date, discharge_diagnosis, patient_id, doctor_id, ward_id 
			  FROM hospitalization WHERE id = $1`
	hospitalization := &models.Hospitalization{}
	err := r.db.QueryRow(query, id).Scan(&hospitalization.ID, &hospitalization.Type, &hospitalization.Status, 
		&hospitalization.BeginningDate, &hospitalization.PrimaryDiagnosis, &hospitalization.DischargeDate, 
		&hospitalization.DischargeDiagnosis, &hospitalization.PatientID, &hospitalization.DoctorID, &hospitalization.WardID)
	if err != nil {
		return nil, err
	}
	return hospitalization, nil
}

func (r *HospitalizationRepository) GetAll() ([]models.Hospitalization, error) {
	query := `SELECT id, type, status, beginning_date, primary_diagnosis, discharge_date, discharge_diagnosis, patient_id, doctor_id, ward_id FROM hospitalization`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hospitalizations []models.Hospitalization
	for rows.Next() {
		var hospitalization models.Hospitalization
		err := rows.Scan(&hospitalization.ID, &hospitalization.Type, &hospitalization.Status,
			&hospitalization.BeginningDate, &hospitalization.PrimaryDiagnosis, &hospitalization.DischargeDate,
			&hospitalization.DischargeDiagnosis, &hospitalization.PatientID, &hospitalization.DoctorID, &hospitalization.WardID)
		if err != nil {
			return nil, err
		}
		hospitalizations = append(hospitalizations, hospitalization)
	}
	return hospitalizations, nil
}

func (r *HospitalizationRepository) Update(hospitalization *models.Hospitalization) error {
	query := `UPDATE hospitalization SET type=$1, status=$2, beginning_date=$3, primary_diagnosis=$4, 
			  discharge_date=$5, discharge_diagnosis=$6, patient_id=$7, doctor_id=$8, ward_id=$9 WHERE id=$10`
	_, err := r.db.Exec(query, hospitalization.Type, hospitalization.Status, hospitalization.BeginningDate,
		hospitalization.PrimaryDiagnosis, hospitalization.DischargeDate, hospitalization.DischargeDiagnosis,
		hospitalization.PatientID, hospitalization.DoctorID, hospitalization.WardID, hospitalization.ID)
	return err
}

func (r *HospitalizationRepository) Delete(id int64) error {
	query := `DELETE FROM hospitalization WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
