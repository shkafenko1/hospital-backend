package service

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/repository"
)

type DoctorService struct {
	repo *repository.DoctorRepository
}

func NewDoctorService(repo *repository.DoctorRepository) *DoctorService {
	return &DoctorService{repo: repo}
}

func (s *DoctorService) Create(doctor *models.Doctor) error {
	return s.repo.Create(doctor)
}

func (s *DoctorService) GetByID(id int64) (*models.Doctor, error) {
	return s.repo.GetByID(id)
}

func (s *DoctorService) GetAll() ([]models.Doctor, error) {
	return s.repo.GetAll()
}

func (s *DoctorService) Update(doctor *models.Doctor) error {
	return s.repo.Update(doctor)
}

func (s *DoctorService) Delete(id int64) error {
	return s.repo.Delete(id)
}
