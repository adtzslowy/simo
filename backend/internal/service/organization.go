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

func (s *OrganizationService) Create(
	ctx context.Context,
	req model.CreateOrganizationRequest,
) (*model.Organization, error) {
	organization := &model.Organization{
		OrganizationTypeID:   req.OrganizationTypeID,
		InstitutionID:        req.InstitutionID,
		AcademicDepartmentID: req.AcademicDepartmentID,
		Name:                 req.Name,
		Code:                 req.Code,
		Description:          req.Description,
		LogoURL:              req.LogoURL,
		IsActive:             req.IsActive,
	}

	if err := s.repository.Create(ctx, organization); err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *OrganizationService) Update(
	ctx context.Context,
	id uuid.UUID,
	req model.UpdateOrganizationRequest,
) (*model.Organization, error) {
	organization := &model.Organization{
		ID:                   id,
		OrganizationTypeID:   req.OrganizationTypeID,
		InstitutionID:        req.InstitutionID,
		AcademicDepartmentID: req.AcademicDepartmentID,
		Name:                 req.Name,
		Code:                 req.Code,
		Description:          req.Description,
		LogoURL:              req.LogoURL,
		IsActive:             req.IsActive,
	}

	if err := s.repository.Update(ctx, organization); err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *OrganizationService) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return s.repository.Delete(ctx, id)
}
