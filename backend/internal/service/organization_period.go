package service

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/google/uuid"
)

type OrganizationPeriodService struct {
	repository *repository.OrganizationPeriodRepository
}

func NewOrganizationPeriodService(
	repository *repository.OrganizationPeriodRepository,
) *OrganizationPeriodService {
	return &OrganizationPeriodService{
		repository: repository,
	}
}

func (s *OrganizationPeriodService) GetByOrganizationID(
	ctx context.Context,
	organizationID uuid.UUID,
) ([]model.OrganizationPeriod, error) {
	return s.repository.FindByOrganizationID(
		ctx, organizationID,
	)
}

func (s *OrganizationPeriodService) GetByID(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) (*model.OrganizationPeriod, error) {
	return s.repository.FindByID(
		ctx,
		organizationID,
		id,
	)
}

func (s *OrganizationPeriodService) Create(
	ctx context.Context,
	organizationID uuid.UUID,
	req model.CreateOrganizationPeriodRequest,
) (*model.OrganizationPeriod, error) {
	period := &model.OrganizationPeriod{
		OrganizationID: organizationID,
		Name:           req.Name,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		IsActive:       req.IsActive,
	}

	if err := s.repository.Create(ctx, period); err != nil {
		return nil, err
	}

	return period, nil
}

func (s *OrganizationPeriodService) Update(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
	req model.UpdateOrganizationPeriodRequest,
) (*model.OrganizationPeriod, error) {
	period := &model.OrganizationPeriod{
		ID:             id,
		OrganizationID: organizationID,
		Name:           req.Name,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		IsActive:       req.IsActive,
	}

	if err := s.repository.Update(ctx, period); err != nil {
		return nil, err
	}

	return period, nil
}

func (s *OrganizationPeriodService) Delete(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) error {
	return s.repository.Delete(
		ctx,
		organizationID,
		id,
	)
}
