package repository

import (
	"database/sql"
	"hospital-backend/internal/domain/models"
)

type MedicineRepository struct {
	db *sql.DB
}

func NewMedicineRepository(db *sql.DB) *MedicineRepository {
	return &MedicineRepository{db: db}
}

func (r *MedicineRepository) Create(medicine *models.Medicine) error {
	query := `INSERT INTO medicine (article, name, active_substance, form, contraindications, prescription) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRow(query, medicine.Article, medicine.Name, medicine.ActiveSubstance, 
		medicine.Form, medicine.Contraindications, medicine.Prescription).Scan(&medicine.ID)
}

func (r *MedicineRepository) GetByID(id int64) (*models.Medicine, error) {
	query := `SELECT id, article, name, active_substance, form, contraindications, prescription 
			  FROM medicine WHERE id = $1`
	medicine := &models.Medicine{}
	err := r.db.QueryRow(query, id).Scan(&medicine.ID, &medicine.Article, &medicine.Name, 
		&medicine.ActiveSubstance, &medicine.Form, &medicine.Contraindications, &medicine.Prescription)
	if err != nil {
		return nil, err
	}
	return medicine, nil
}

func (r *MedicineRepository) GetAll() ([]models.Medicine, error) {
	query := `SELECT id, article, name, active_substance, form, contraindications, prescription FROM medicine`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medicines []models.Medicine
	for rows.Next() {
		var medicine models.Medicine
		err := rows.Scan(&medicine.ID, &medicine.Article, &medicine.Name, &medicine.ActiveSubstance,
			&medicine.Form, &medicine.Contraindications, &medicine.Prescription)
		if err != nil {
			return nil, err
		}
		medicines = append(medicines, medicine)
	}
	return medicines, nil
}

func (r *MedicineRepository) Update(medicine *models.Medicine) error {
	query := `UPDATE medicine SET article=$1, name=$2, active_substance=$3, form=$4, 
			  contraindications=$5, prescription=$6 WHERE id=$7`
	_, err := r.db.Exec(query, medicine.Article, medicine.Name, medicine.ActiveSubstance,
		medicine.Form, medicine.Contraindications, medicine.Prescription, medicine.ID)
	return err
}

func (r *MedicineRepository) Delete(id int64) error {
	query := `DELETE FROM medicine WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
