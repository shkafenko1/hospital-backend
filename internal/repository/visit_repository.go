package repository

import (
	"database/sql"
	"hospital-backend/internal/domain/models"
)

type VisitRepository struct {
	db *sql.DB
}

func NewVisitRepository(db *sql.DB) *VisitRepository {
	return &VisitRepository{db: db}
}

func (r *VisitRepository) Create(visit *models.Visit) error {
	query := `INSERT INTO visit (date_time, symptoms, diagnosis, advise, price, patient_id, doctor_id) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	return r.db.QueryRow(query, visit.DateTime, visit.Symptoms, visit.Diagnosis, 
		visit.Advise, visit.Price, visit.PatientID, visit.DoctorID).Scan(&visit.ID)
}

func (r *VisitRepository) GetByID(id int64) (*models.Visit, error) {
	query := `SELECT id, date_time, symptoms, diagnosis, advise, price, patient_id, doctor_id 
			  FROM visit WHERE id = $1`
	visit := &models.Visit{}
	err := r.db.QueryRow(query, id).Scan(&visit.ID, &visit.DateTime, &visit.Symptoms, 
		&visit.Diagnosis, &visit.Advise, &visit.Price, &visit.PatientID, &visit.DoctorID)
	if err != nil {
		return nil, err
	}
	return visit, nil
}

func (r *VisitRepository) GetAll() ([]models.Visit, error) {
	query := `SELECT id, date_time, symptoms, diagnosis, advise, price, patient_id, doctor_id FROM visit`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visits []models.Visit
	for rows.Next() {
		var visit models.Visit
		err := rows.Scan(&visit.ID, &visit.DateTime, &visit.Symptoms, &visit.Diagnosis,
			&visit.Advise, &visit.Price, &visit.PatientID, &visit.DoctorID)
		if err != nil {
			return nil, err
		}
		visits = append(visits, visit)
	}
	return visits, nil
}

func (r *VisitRepository) Update(visit *models.Visit) error {
	query := `UPDATE visit SET date_time=$1, symptoms=$2, diagnosis=$3, advise=$4, 
			  price=$5, patient_id=$6, doctor_id=$7 WHERE id=$8`
	_, err := r.db.Exec(query, visit.DateTime, visit.Symptoms, visit.Diagnosis,
		visit.Advise, visit.Price, visit.PatientID, visit.DoctorID, visit.ID)
	return err
}

func (r *VisitRepository) Delete(id int64) error {
	query := `DELETE FROM visit WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
