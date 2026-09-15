package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/response"
	"github.com/adtzslowy/simo/internal/service"
	"github.com/google/uuid"
)

type OrganizationMemberHandler struct {
	service *service.OrganizationMemberService
}

func NewOrganizationMemberHandler(
	service *service.OrganizationMemberService,
) *OrganizationMemberHandler {
	return &OrganizationMemberHandler{
		service: service,
	}
}

func (h *OrganizationMemberHandler) GetAllByOrganizationID(
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

	members, err := h.service.GetAllByOrganizationID(
		r.Context(),
		organizationID,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to get organization members",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"organization members retrieved successfully",
		members,
	)
}

func (h *OrganizationMemberHandler) GetByID(
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
			"invalid member id",
			nil,
		)
		return
	}

	member, err := h.service.GetByID(
		r.Context(),
		organizationID,
		id,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusNotFound,
			"organization member not found",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"organization member retrieved successfully",
		member,
	)
}

func (h *OrganizationMemberHandler) Create(
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

	var req model.CreateOrganizationMemberRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
		return
	}

	member, err := h.service.Create(
		r.Context(),
		organizationID,
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to create organization member",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusCreated,
		"organization member created successfully",
		member,
	)
}

func (h *OrganizationMemberHandler) Update(
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
			"invalid member id",
			nil,
		)
		return
	}

	var req model.UpdateOrganizationMemberRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
		return
	}

	member, err := h.service.Update(
		r.Context(),
		organizationID,
		id,
		req,
	)
	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			"failed to update organization member",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"organization member updated successfully",
		member,
	)
}

func (h *OrganizationMemberHandler) Delete(
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
			"invalid member id",
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
			"failed to delete organization member",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"organization member deleted successfully",
		nil,
	)
}
