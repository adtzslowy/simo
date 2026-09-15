package model

import "github.com/google/uuid"

type CreateDepartmentRequest struct {
	PeriodID    *uuid.UUID `json:"period_id,omitempty"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
	Name        string     `json:"name"`
	Code        string     `json:"code"`
	Description *string    `json:"description,omitempty"`
	IsActive    bool       `json:"is_active"`
}

type UpdateDepartmentRequest struct {
	PeriodID    *uuid.UUID `json:"period_id,omitempty"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty"`
	Name        string     `json:"name"`
	Code        string     `json:"code"`
	Description *string    `json:"description,omitempty"`
	IsActive    bool       `json:"is_active"`
}
