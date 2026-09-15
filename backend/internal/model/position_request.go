package model

type CreatePositionRequest struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Description *string `json:"description,omitempty"`
	IsActive    bool    `json:"is_active"`
}

type UpdatePositionRequest struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Description *string `json:"description,omitempty"`
	IsActive    bool    `json:"is_active"`
}
