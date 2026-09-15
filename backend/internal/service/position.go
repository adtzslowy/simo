package service

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/google/uuid"
)

type PositionService struct {
	repository *repository.PositionRepository
}

func NewPositionService(
	repository *repository.PositionRepository,
) *PositionService {
	return &PositionService{
		repository: repository,
	}
}

func (s *PositionService) GetAllByOrganizationID(
	ctx context.Context,
	organizationID uuid.UUID,
) ([]model.Position, error) {
	return s.repository.FindAllByOrganizationID(
		ctx,
		organizationID,
	)
}

func (s *PositionService) GetByID(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) (*model.Position, error) {
	return s.repository.FindByID(
		ctx,
		organizationID,
		id,
	)
}

func (s *PositionService) Create(
	ctx context.Context,
	organizationID uuid.UUID,
	req model.CreatePositionRequest,
) (*model.Position, error) {
	position := &model.Position{
		OrganizationID: organizationID,
		Name:           req.Name,
		Code:           req.Code,
		Description:    req.Description,
		IsActive:       req.IsActive,
	}

	if err := s.repository.Create(
		ctx,
		position,
	); err != nil {
		return nil, err
	}

	return position, nil
}

func (s *PositionService) Update(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
	req model.UpdatePositionRequest,
) (*model.Position, error) {
	position := &model.Position{
		ID:             id,
		OrganizationID: organizationID,
		Name:           req.Name,
		Code:           req.Code,
		Description:    req.Description,
		IsActive:       req.IsActive,
	}

	if err := s.repository.Update(
		ctx,
		position,
	); err != nil {
		return nil, err
	}

	return position, nil
}

func (s *PositionService) Delete(
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
