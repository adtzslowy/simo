package handler

import (
	"net/http"

	"github.com/adtzslowy/simo/internal/response"
	"github.com/adtzslowy/simo/internal/service"
	"github.com/google/uuid"
)

type OrganizationHandler struct {
	service *service.OrganizationService
}

func NewOrganizationHandler(
	service *service.OrganizationService,
) *OrganizationHandler {
	return &OrganizationHandler{
		service: service,
	}
}

func (h *OrganizationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	organization, err := h.service.GetAll(r.Context())
	if err != nil {
		response.Error(
			w, http.StatusInternalServerError, "failed to get organization", nil,
		)
		return
	}

	response.Success(
		w, http.StatusOK, "organization retrieve successfully", organization,
	)
}

func (h *OrganizationHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(
			w, http.StatusBadRequest, "failed to get organization data", nil,
		)
		return
	}

	organization, err := h.service.GetByID(
		r.Context(),
		id,
	)
	if err != nil {
		response.Error(
			w, http.StatusBadRequest, "organization not found", nil,
		)
		return
	}

	response.Success(
		w, http.StatusOK, "organization retrieve successfully", organization,
	)
}
