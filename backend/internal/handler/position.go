package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/response"
	"github.com/adtzslowy/simo/internal/service"
	"github.com/google/uuid"
)

type PositionHandler struct {
	service *service.PositionService
}

func NewPositionHandler(
	service *service.PositionService,
) *PositionHandler {
	return &PositionHandler{
		service: service,
	}
}

func (h *PositionHandler) GetAllByOrganizationID(
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

	positions, err := h.service.GetAllByOrganizationID(
		r.Context(),
		organizationID,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to get positions",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"positions retrieved successfully",
		positions,
	)
}

func (h *PositionHandler) GetByID(
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
			"invalid position id",
			nil,
		)
		return
	}

	position, err := h.service.GetByID(
		r.Context(),
		organizationID,
		id,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusNotFound,
			"position not found",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"position retrieved successfully",
		position,
	)
}

func (h *PositionHandler) Create(
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

	var req model.CreatePositionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
		return
	}

	position, err := h.service.Create(
		r.Context(),
		organizationID,
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to create position",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusCreated,
		"position created successfully",
		position,
	)
}

func (h *PositionHandler) Update(
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
			"invalid position id",
			nil,
		)
		return
	}

	var req model.UpdatePositionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
		return
	}

	position, err := h.service.Update(
		r.Context(),
		organizationID,
		id,
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to update position",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"position updated successfully",
		position,
	)
}

func (h *PositionHandler) Delete(
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
			"invalid position id",
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
			"failed to delete position",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"position deleted successfully",
		nil,
	)
}
