package service

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/repository"
)

type MedicineService struct {
	repo *repository.MedicineRepository
}

func NewMedicineService(repo *repository.MedicineRepository) *MedicineService {
	return &MedicineService{repo: repo}
}

func (s *MedicineService) Create(medicine *models.Medicine) error {
	return s.repo.Create(medicine)
}

func (s *MedicineService) GetByID(id int64) (*models.Medicine, error) {
	return s.repo.GetByID(id)
}

func (s *MedicineService) GetAll() ([]models.Medicine, error) {
	return s.repo.GetAll()
}

func (s *MedicineService) Update(medicine *models.Medicine) error {
	return s.repo.Update(medicine)
}

func (s *MedicineService) Delete(id int64) error {
	return s.repo.Delete(id)
}
