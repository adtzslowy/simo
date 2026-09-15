package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/response"
	"github.com/adtzslowy/simo/internal/service"
	"github.com/google/uuid"
)

type DepartmentHandler struct {
	service *service.DepartmentService
}

func NewDepartmentHandler(
	service *service.DepartmentService,
) *DepartmentHandler {
	return &DepartmentHandler{
		service: service,
	}
}

func (h *DepartmentHandler) GetAllByOrganizationID(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, err := uuid.Parse(
		r.PathValue("organization_id"),
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization id",
			nil,
		)
		return
	}

	departments, err := h.service.GetAllByOrganizationID(
		r.Context(),
		organizationID,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to get departments",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"departments retrieved successfully",
		departments,
	)
}

func (h *DepartmentHandler) GetByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, err := uuid.Parse(
		r.PathValue("organization_id"),
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization id",
			nil,
		)
		return
	}

	id, err := uuid.Parse(
		r.PathValue("id"),
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid department id",
			nil,
		)
		return
	}

	department, err := h.service.GetByID(
		r.Context(),
		organizationID,
		id,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusNotFound,
			"department not found",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"department retrieved successfully",
		department,
	)
}

func (h *DepartmentHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, err := uuid.Parse(
		r.PathValue("organization_id"),
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization id",
			nil,
		)
		return
	}

	var req model.CreateDepartmentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
		return
	}

	department, err := h.service.Create(
		r.Context(),
		organizationID,
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to create department",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusCreated,
		"department created successfully",
		department,
	)
}

func (h *DepartmentHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, err := uuid.Parse(
		r.PathValue("organization_id"),
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization id",
			nil,
		)
		return
	}

	id, err := uuid.Parse(
		r.PathValue("id"),
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid department id",
			nil,
		)
		return
	}

	var req model.UpdateDepartmentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
		return
	}

	department, err := h.service.Update(
		r.Context(),
		organizationID,
		id,
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to update department",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"department updated successfully",
		department,
	)
}

func (h *DepartmentHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, err := uuid.Parse(
		r.PathValue("organization_id"),
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization id",
			nil,
		)
		return
	}

	id, err := uuid.Parse(
		r.PathValue("id"),
	)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid department id",
			nil,
		)
		return
	}

	if err := h.service.Delete(
		r.Context(),
		organizationID,
		id,
	); err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to delete department",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"department deleted successfully",
		nil,
	)
}
