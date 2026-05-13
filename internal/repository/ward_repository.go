package repository

import (
	"database/sql"
	"hospital-backend/internal/domain/models"
)

type WardRepository struct {
	db *sql.DB
}

func NewWardRepository(db *sql.DB) *WardRepository {
	return &WardRepository{db: db}
}

func (r *WardRepository) Create(ward *models.Ward) error {
	query := `INSERT INTO ward ("number", floor, capacity, price, comfort, busy) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRow(query, ward.Number, ward.Floor, ward.Capacity, ward.Price, ward.Comfort, ward.Busy).Scan(&ward.ID)
}

func (r *WardRepository) GetByID(id int64) (*models.Ward, error) {
	query := `SELECT id, "number", floor, capacity, price, comfort, busy FROM ward WHERE id = $1`
	ward := &models.Ward{}
	err := r.db.QueryRow(query, id).Scan(&ward.ID, &ward.Number, &ward.Floor, &ward.Capacity, 
		&ward.Price, &ward.Comfort, &ward.Busy)
	if err != nil {
		return nil, err
	}
	return ward, nil
}

func (r *WardRepository) GetAll() ([]models.Ward, error) {
	query := `SELECT id, "number", floor, capacity, price, comfort, busy FROM ward`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wards []models.Ward
	for rows.Next() {
		var ward models.Ward
		err := rows.Scan(&ward.ID, &ward.Number, &ward.Floor, &ward.Capacity, 
			&ward.Price, &ward.Comfort, &ward.Busy)
		if err != nil {
			return nil, err
		}
		wards = append(wards, ward)
	}
	return wards, nil
}

func (r *WardRepository) Update(ward *models.Ward) error {
	query := `UPDATE ward SET "number"=$1, floor=$2, capacity=$3, price=$4, comfort=$5, busy=$6 WHERE id=$7`
	_, err := r.db.Exec(query, ward.Number, ward.Floor, ward.Capacity, ward.Price, ward.Comfort, ward.Busy, ward.ID)
	return err
}

func (r *WardRepository) Delete(id int64) error {
	query := `DELETE FROM ward WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
