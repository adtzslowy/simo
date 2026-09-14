package service

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/google/uuid"
)

type OrganizationService struct {
	repository *repository.OrganizationRepository
}

func NewOrganizationService(repository *repository.OrganizationRepository) *OrganizationService {
	return &OrganizationService{
		repository: repository,
	}
}

func (s *OrganizationService) GetAll(ctx context.Context) ([]model.Organization, error) {
	return s.repository.FindAll(ctx)
}

func (s *OrganizationService) GetByID(ctx context.Context, id uuid.UUID) (*model.Organization, error) {
	return s.repository.FindByID(ctx, id)
}
