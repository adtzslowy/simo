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
