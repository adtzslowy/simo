package model

import (
	"time"

	"github.com/google/uuid"
)

type MemberPosition struct {
	ID           uuid.UUID  `json:"id"`
	MemberID     uuid.UUID  `json:"member_id"`
	PositionID   uuid.UUID  `json:"position_id"`
	DepartmentID *uuid.UUID `json:"department_id,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
