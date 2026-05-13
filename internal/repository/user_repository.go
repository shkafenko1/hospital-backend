package repository

import (
	"database/sql"
	"hospital-backend/internal/domain/storage"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (*storage.UserDTO, error) {
	dto := &storage.UserDTO{}
	query := `SELECT id, username, password_hash, role, created_at FROM users WHERE username = $1`
	err := r.db.QueryRow(query, username).Scan(&dto.ID, &dto.Username, &dto.PasswordHash, &dto.Role, &dto.CreatedAt)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (r *UserRepository) GetByID(id int64) (*storage.UserDTO, error) {
	dto := &storage.UserDTO{}
	query := `SELECT id, username, password_hash, role, created_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&dto.ID, &dto.Username, &dto.PasswordHash, &dto.Role, &dto.CreatedAt)
	if err != nil {
		return nil, err
	}
	return dto, nil
}
