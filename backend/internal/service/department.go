package service

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/google/uuid"
)

type DepartmentService struct {
	repository *repository.DepartmentRepository
}

func NewDepartmentService(
	repository *repository.DepartmentRepository,
) *DepartmentService {
	return &DepartmentService{
		repository: repository,
	}
}

func (s *DepartmentService) GetAllByOrganizationID(
	ctx context.Context,
	organizationID uuid.UUID,
) ([]model.Department, error) {
	return s.repository.FindAllByOrganizationID(
		ctx,
		organizationID,
	)
}

func (s *DepartmentService) GetByID(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) (*model.Department, error) {
	return s.repository.FindByID(
		ctx,
		organizationID,
		id,
	)
}

func (s *DepartmentService) Create(
	ctx context.Context,
	organizationID uuid.UUID,
	req model.CreateDepartmentRequest,
) (*model.Department, error) {
	department := &model.Department{
		OrganizationID: organizationID,
		PeriodID:       req.PeriodID,
		ParentID:       req.ParentID,
		Name:           req.Name,
		Code:           req.Code,
		Description:    req.Description,
		IsActive:       req.IsActive,
	}

	if err := s.repository.Create(
		ctx,
		department,
	); err != nil {
		return nil, err
	}

	return department, nil
}

func (s *DepartmentService) Update(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
	req model.UpdateDepartmentRequest,
) (*model.Department, error) {
	department := &model.Department{
		ID:             id,
		OrganizationID: organizationID,
		PeriodID:       req.PeriodID,
		ParentID:       req.ParentID,
		Name:           req.Name,
		Code:           req.Code,
		Description:    req.Description,
		IsActive:       req.IsActive,
	}

	if err := s.repository.Update(
		ctx,
		department,
	); err != nil {
		return nil, err
	}

	return department, nil
}

func (s *DepartmentService) Delete(
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
