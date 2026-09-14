package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adtzslowy/simo/internal/model"
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

func (h *OrganizationTypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateOrganizationTypeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w, http.StatusBadRequest, "invalid body request", nil,
		)
		return
	}

	organizationType, err := h.service.Create(
		r.Context(),
		req,
	)

	if err != nil {
		response.Error(
			w, http.StatusInternalServerError, "failed to create organization type", nil,
		)
		return
	}

	response.Success(
		w, http.StatusCreated, "organization type created successfully", organizationType,
	)
}

func (h *OrganizationTypeHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(
			w, http.StatusBadRequest, "invalid organization type id", nil,
		)
		return
	}

	var req model.UpdateOrganizationTypeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w, http.StatusBadRequest, "invalid request body", nil,
		)
		return
	}

	organizationType, err := h.service.Update(
		r.Context(),
		id,
		req,
	)

	if err != nil {
		response.Error(
			w, http.StatusInternalServerError, "failed to update organization type", nil,
		)
		return
	}

	response.Success(
		w, http.StatusOK, "organization type updated successfully", organizationType,
	)
}

func (h *OrganizationTypeHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(
			w, http.StatusBadRequest, "invalid organization type id", nil,
		)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to delete organization type",
			nil,
		)
		return
	}

	response.Success(
		w, http.StatusOK, "organization type deleted successfully", nil,
	)
}
