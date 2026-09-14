package handler

import (
	"net/http"

	"github.com/adtzslowy/simo/internal/response"
	"github.com/adtzslowy/simo/internal/service"
	"github.com/google/uuid"
)

type OrganizationTypeHandler struct {
	service *service.OrganizationTypeService
}

func NewOrganizationTypeHandler(
	service *service.OrganizationTypeService,
) *OrganizationTypeHandler {
	return &OrganizationTypeHandler{
		service: service,
	}
}

func (h *OrganizationTypeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	types, err := h.service.GetAll(r.Context())
	if err != nil {
		response.Error(
			w, http.StatusInternalServerError, "failed to get organization type", nil,
		)
		return
	}

	response.Success(
		w, http.StatusOK, "organization retrieve successfully", types,
	)
}

func (h *OrganizationTypeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(
			w, http.StatusBadRequest, "failed to get organization data", nil,
		)
		return
	}

	types, err := h.service.GetByID(
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
		w, http.StatusOK, "organization type retrieve successfully", types,
	)
}
