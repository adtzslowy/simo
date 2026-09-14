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

func (s *OrganizationTypeService) Create(
	ctx context.Context,
	req model.CreateOrganizationTypeRequest,
) (*model.OrganizationType, error) {
	organizationType := &model.OrganizationType{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	}

	if err := s.repository.Create(ctx, organizationType); err != nil {
		return nil, err
	}

	return organizationType, nil
}

func (s *OrganizationTypeService) Update(
	ctx context.Context,
	id uuid.UUID,
	req model.UpdateOrganizationTypeRequest,
) (*model.OrganizationType, error) {
	organizationType := &model.OrganizationType{
		ID:          id,
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	}

	if err := s.repository.Update(ctx, organizationType); err != nil {
		return nil, err
	}

	return organizationType, nil
}

func (s *OrganizationTypeService) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return s.repository.Delete(ctx, id)
}
