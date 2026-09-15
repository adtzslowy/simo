package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/response"
	"github.com/adtzslowy/simo/internal/service"
	"github.com/google/uuid"
)

type MemberPositionHandler struct {
	service *service.MemberPositionService
}

func NewMemberPositionHandler(
	service *service.MemberPositionService,
) *MemberPositionHandler {
	return &MemberPositionHandler{
		service: service,
	}
}

func parseMemberPositionIDs(
	r *http.Request,
) (uuid.UUID, uuid.UUID, error) {
	organizationID, err := uuid.Parse(
		r.PathValue("organization_id"),
	)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	memberID, err := uuid.Parse(
		r.PathValue("member_id"),
	)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	return organizationID, memberID, nil
}

func (h *MemberPositionHandler) GetAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, memberID, err := parseMemberPositionIDs(r)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization or member id",
			nil,
		)
		return
	}

	assignments, err := h.service.GetAllByMemberID(
		r.Context(),
		organizationID,
		memberID,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to get member positions",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"member positions retrieved successfully",
		assignments,
	)
}

func (h *MemberPositionHandler) GetByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, memberID, err := parseMemberPositionIDs(r)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization or member id",
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
			"invalid assignment id",
			nil,
		)
		return
	}

	assignment, err := h.service.GetByID(
		r.Context(),
		organizationID,
		memberID,
		id,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusNotFound,
			"member position not found",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"member position retrieved successfully",
		assignment,
	)
}

func (h *MemberPositionHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, memberID, err := parseMemberPositionIDs(r)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization or member id",
			nil,
		)
		return
	}

	var req model.CreateMemberPositionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
		return
	}

	assignment, err := h.service.Create(
		r.Context(),
		organizationID,
		memberID,
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to create member position",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusCreated,
		"member position created successfully",
		assignment,
	)
}

func (h *MemberPositionHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, memberID, err := parseMemberPositionIDs(r)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization or member id",
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
			"invalid assignment id",
			nil,
		)
		return
	}

	var req model.UpdateMemberPositionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
		return
	}

	assignment, err := h.service.Update(
		r.Context(),
		organizationID,
		memberID,
		id,
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to update member position",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"member position updated successfully",
		assignment,
	)
}

func (h *MemberPositionHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationID, memberID, err := parseMemberPositionIDs(r)
	if err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid organization or member id",
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
			"invalid assignment id",
			nil,
		)
		return
	}

	if err := h.service.Delete(
		r.Context(),
		organizationID,
		memberID,
		id,
	); err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to delete member position",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"member position deleted successfully",
		nil,
	)
}
