package service

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/google/uuid"
)

type OrganizationTypeService struct {
	repository *repository.OrganizationTypeRepository
}

func NewOrganizationTypeService(
	repository *repository.OrganizationTypeRepository,
) *OrganizationTypeService {
	return &OrganizationTypeService{
		repository: repository,
	}
}

func (s *OrganizationTypeService) GetAll(
	ctx context.Context,
) ([]model.OrganizationType, error) {
	return s.repository.FindAll(ctx)
}

func (s *OrganizationTypeService) GetByID(ctx context.Context, id uuid.UUID) (*model.OrganizationType, error) {
	return s.repository.FindByID(ctx, id)
}
