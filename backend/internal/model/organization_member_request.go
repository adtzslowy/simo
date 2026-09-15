package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateOrganizationMemberRequest struct {
	UserID   uuid.UUID  `json:"user_id"`
	PeriodID *uuid.UUID `json:"period_id,omitempty"`
	Status   string     `json:"status"`
	JoinedAt *time.Time `json:"joined_at,omitempty"`
	LeftAt   *time.Time `json:"left_at,omitempty"`
}

type UpdateOrganizationMemberRequest struct {
	PeriodID *uuid.UUID `json:"period_id,omitempty"`
	Status   string     `json:"status"`
	JoinedAt *time.Time `json:"joined_at,omitempty"`
	LeftAt   *time.Time `json:"left_at,omitempty"`
}
