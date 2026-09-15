package service

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/google/uuid"
)

type OrganizationMemberService struct {
	repository *repository.OrganizationMemberRepository
}

func NewOrganizationMemberService(
	repository *repository.OrganizationMemberRepository,
) *OrganizationMemberService {
	return &OrganizationMemberService{
		repository: repository,
	}
}

func (s *OrganizationMemberService) GetAllByOrganizationID(
	ctx context.Context,
	organizationID uuid.UUID,
) ([]model.OrganizationMember, error) {
	return s.repository.FindAllByOrganizationID(
		ctx,
		organizationID,
	)
}

func (s *OrganizationMemberService) GetByID(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) (*model.OrganizationMember, error) {
	return s.repository.FindByID(
		ctx,
		organizationID,
		id,
	)
}

func (s *OrganizationMemberService) Create(
	ctx context.Context,
	organizationID uuid.UUID,
	req model.CreateOrganizationMemberRequest,
) (*model.OrganizationMember, error) {
	member := &model.OrganizationMember{
		OrganizationID: organizationID,
		UserID:         req.UserID,
		PeriodID:       req.PeriodID,
		Status:         req.Status,
		JoinedAt:       req.JoinedAt,
		LeftAt:         req.LeftAt,
	}

	if err := s.repository.Create(
		ctx,
		member,
	); err != nil {
		return nil, err
	}

	return member, nil
}

func (s *OrganizationMemberService) Update(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
	req model.UpdateOrganizationMemberRequest,
) (*model.OrganizationMember, error) {
	member := &model.OrganizationMember{
		ID:             id,
		OrganizationID: organizationID,
		PeriodID:       req.PeriodID,
		Status:         req.Status,
		JoinedAt:       req.JoinedAt,
		LeftAt:         req.LeftAt,
	}

	if err := s.repository.Update(
		ctx,
		member,
	); err != nil {
		return nil, err
	}

	return member, nil
}

func (s *OrganizationMemberService) Delete(
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
