package model

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID                   uuid.UUID  `json:"id"`
	OrganizationTypeID   uuid.UUID  `json:"organization_type_id"`
	InstitutionID        *uuid.UUID `json:"institution_id,omitempty"`
	AcademicDepartmentID *uuid.UUID `json:"academic_department_id,omitempty"`
	Name                 string     `json:"name"`
	Code                 string     `json:"code"`
	Description          *string    `json:"description,omitempty"`
	LogoURL              *string    `json:"logo_url,omitempty"`
	IsActive             bool       `json:"is_active"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}
