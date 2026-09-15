package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adtzslowy/simo/internal/model"
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

func (h *OrganizationHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req model.CreateOrganizationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
		return
	}

	organization, err := h.service.Create(
		r.Context(),
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to create organization",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusCreated,
		"organization created successfully",
		organization,
	)
}

func (h *OrganizationHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization id",
			nil,
		)
		return
	}

	var req model.UpdateOrganizationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
		return
	}

	organization, err := h.service.Update(
		r.Context(),
		id,
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to update organization",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"organization updated successfully",
		organization,
	)
}

func (h *OrganizationHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization id",
			nil,
		)
		return
	}

	if err := h.service.Delete(
		r.Context(),
		id,
	); err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to delete organization",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"organization deleted successfully",
		nil,
	)
}
