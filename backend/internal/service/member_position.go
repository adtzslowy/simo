package service

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/google/uuid"
)

type MemberPositionService struct {
	repository *repository.MemberPositionRepository
}

func NewMemberPositionService(
	repository *repository.MemberPositionRepository,
) *MemberPositionService {
	return &MemberPositionService{
		repository: repository,
	}
}

func (s *MemberPositionService) GetAllByMemberID(
	ctx context.Context,
	organizationID uuid.UUID,
	memberID uuid.UUID,
) ([]model.MemberPosition, error) {
	return s.repository.FindAllByMemberID(
		ctx,
		organizationID,
		memberID,
	)
}

func (s *MemberPositionService) GetByID(
	ctx context.Context,
	organizationID uuid.UUID,
	memberID uuid.UUID,
	id uuid.UUID,
) (*model.MemberPosition, error) {
	return s.repository.FindByID(
		ctx,
		organizationID,
		memberID,
		id,
	)
}

func (s *MemberPositionService) Create(
	ctx context.Context,
	organizationID uuid.UUID,
	memberID uuid.UUID,
	req model.CreateMemberPositionRequest,
) (*model.MemberPosition, error) {
	assignment := &model.MemberPosition{
		MemberID:     memberID,
		PositionID:   req.PositionID,
		DepartmentID: req.DepartmentID,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		IsActive:     req.IsActive,
	}

	if err := s.repository.Create(
		ctx,
		organizationID,
		memberID,
		assignment,
	); err != nil {
		return nil, err
	}

	return assignment, nil
}

func (s *MemberPositionService) Update(
	ctx context.Context,
	organizationID uuid.UUID,
	memberID uuid.UUID,
	id uuid.UUID,
	req model.UpdateMemberPositionRequest,
) (*model.MemberPosition, error) {
	assignment := &model.MemberPosition{
		ID:           id,
		MemberID:     memberID,
		PositionID:   req.PositionID,
		DepartmentID: req.DepartmentID,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		IsActive:     req.IsActive,
	}

	if err := s.repository.Update(
		ctx,
		organizationID,
		memberID,
		id,
		assignment,
	); err != nil {
		return nil, err
	}

	return assignment, nil
}

func (s *MemberPositionService) Delete(
	ctx context.Context,
	organizationID uuid.UUID,
	memberID uuid.UUID,
	id uuid.UUID,
) error {
	return s.repository.Delete(
		ctx,
		organizationID,
		memberID,
		id,
	)
}
