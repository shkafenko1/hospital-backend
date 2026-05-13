package service

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/repository"
)

type HospitalizationService struct {
	repo *repository.HospitalizationRepository
}

func NewHospitalizationService(repo *repository.HospitalizationRepository) *HospitalizationService {
	return &HospitalizationService{repo: repo}
}

func (s *HospitalizationService) Create(hospitalization *models.Hospitalization) error {
	return s.repo.Create(hospitalization)
}

func (s *HospitalizationService) GetByID(id int64) (*models.Hospitalization, error) {
	return s.repo.GetByID(id)
}

func (s *HospitalizationService) GetAll() ([]models.Hospitalization, error) {
	return s.repo.GetAll()
}

func (s *HospitalizationService) Update(hospitalization *models.Hospitalization) error {
	return s.repo.Update(hospitalization)
}

func (s *HospitalizationService) Delete(id int64) error {
	return s.repo.Delete(id)
}
