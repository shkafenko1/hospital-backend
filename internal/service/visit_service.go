package service

import (
	"hospital-backend/internal/domain/models"
	"hospital-backend/internal/repository"
)

type VisitService struct {
	repo *repository.VisitRepository
}

func NewVisitService(repo *repository.VisitRepository) *VisitService {
	return &VisitService{repo: repo}
}

func (s *VisitService) Create(visit *models.Visit) error {
	return s.repo.Create(visit)
}

func (s *VisitService) GetByID(id int64) (*models.Visit, error) {
	return s.repo.GetByID(id)
}

func (s *VisitService) GetAll() ([]models.Visit, error) {
	return s.repo.GetAll()
}

func (s *VisitService) Update(visit *models.Visit) error {
	return s.repo.Update(visit)
}

func (s *VisitService) Delete(id int64) error {
	return s.repo.Delete(id)
}
