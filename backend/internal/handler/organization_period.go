package handler

import (
	"net/http"

	"github.com/adtzslowy/simo/internal/response"
	"github.com/adtzslowy/simo/internal/service"
	"github.com/google/uuid"
)

type OrganizationPeriodHandler struct {
	service *service.OrganizationPeriodService
}

func NewOrganizationPeriodHandler(
	service *service.OrganizationPeriodService,
) *OrganizationPeriodHandler {
	return &OrganizationPeriodHandler{
		service: service,
	}
}

func (h *OrganizationPeriodHandler) GetByOrganizationID(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, err := uuid.Parse(
		r.PathValue("organization_id"),
	)
	if err != nil {
		response.Error(
			w, http.StatusBadRequest, "invalid organization id", nil,
		)
		return
	}

	periods, err := h.service.GetByOrganizationID(
		r.Context(),
		organizationID,
	)
	if err != nil {
		response.Error(
			w, http.StatusBadRequest, "failed to get organization periods", nil,
		)
		return
	}

	response.Success(
		w, http.StatusOK, "organization periods retrieved successfully", periods,
	)
}
