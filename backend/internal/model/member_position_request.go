package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateMemberPositionRequest struct {
	PositionID   uuid.UUID  `json:"position_id"`
	DepartmentID *uuid.UUID `json:"department_id,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	IsActive     bool       `json:"is_active"`
}

type UpdateMemberPositionRequest struct {
	PositionID   uuid.UUID  `json:"position_id"`
	DepartmentID *uuid.UUID `json:"department_id,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	IsActive     bool       `json:"is_active"`
}
