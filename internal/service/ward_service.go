package service

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/repository"
)

type WardService struct {
	repo *repository.WardRepository
}

func NewWardService(repo *repository.WardRepository) *WardService {
	return &WardService{repo: repo}
}

func (s *WardService) Create(ward *models.Ward) error {
	return s.repo.Create(ward)
}

func (s *WardService) GetByID(id int64) (*models.Ward, error) {
	return s.repo.GetByID(id)
}

func (s *WardService) GetAll() ([]models.Ward, error) {
	return s.repo.GetAll()
}

func (s *WardService) Update(ward *models.Ward) error {
	return s.repo.Update(ward)
}

func (s *WardService) Delete(id int64) error {
	return s.repo.Delete(id)
}
