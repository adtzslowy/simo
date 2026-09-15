package model

import "github.com/google/uuid"

type CreateOrganizationRequest struct {
	OrganizationTypeID   uuid.UUID  `json:"organization_type_id"`
	InstitutionID        *uuid.UUID `json:"institution_id,omitempty"`
	AcademicDepartmentID *uuid.UUID `json:"academic_department_id,omitempty"`
	Name                 string     `json:"name"`
	Code                 string     `json:"code"`
	Description          *string    `json:"description,omitempty"`
	LogoURL              *string    `json:"logo_url,omitempty"`
	IsActive             bool       `json:"is_active"`
}

type UpdateOrganizationRequest struct {
	OrganizationTypeID   uuid.UUID  `json:"organization_type_id"`
	InstitutionID        *uuid.UUID `json:"institution_id,omitempty"`
	AcademicDepartmentID *uuid.UUID `json:"academic_department_id,omitempty"`
	Name                 string     `json:"name"`
	Code                 string     `json:"code"`
	Description          *string    `json:"description,omitempty"`
	LogoURL              *string    `json:"logo_url,omitempty"`
	IsActive             bool       `json:"is_active"`
}
