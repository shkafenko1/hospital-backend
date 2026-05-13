package service

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/repository"
)

type PatientService struct {
	repo *repository.PatientRepository
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func (s *PatientService) Create(patient *models.Patient) error {
	return s.repo.Create(patient)
}

func (s *PatientService) GetByID(id int64) (*models.Patient, error) {
	return s.repo.GetByID(id)
}

func (s *PatientService) GetAll() ([]models.Patient, error) {
	return s.repo.GetAll()
}

func (s *PatientService) Update(patient *models.Patient) error {
	return s.repo.Update(patient)
}

func (s *PatientService) Delete(id int64) error {
	return s.repo.Delete(id)
}
