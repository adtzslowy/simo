package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	appcontext "github.com/adtzslowy/simo/internal/context"
	"github.com/adtzslowy/simo/internal/dto"
	"github.com/adtzslowy/simo/internal/response"
	"github.com/adtzslowy/simo/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(
	service *service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Error(
			w, http.StatusBadRequest, "invalid request body", nil,
		)
		return
	}

	token, err := h.service.Login(
		r.Context(),
		request.Email,
		request.Password,
	)
	if err != nil {
		if errors.Is(err, service.ErrUserInactive) {
			response.Error(
				w, http.StatusForbidden, "user is inactive", nil,
			)
			return
		}

		response.Error(
			w, http.StatusUnauthorized, "invalid email or password.", nil,
		)
		return
	}

	response.Success(
		w, http.StatusOK, "login successfully", dto.LoginResponse{
			AccessToken: token,
			TokenType:   "Bearer",
			ExpiresIn:   86400,
		},
	)
}

func (h *AuthHandler) Me(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, ok := appcontext.GetUserID(r.Context())

	if !ok {
		response.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
			nil,
		)
		return
	}

	user, err := h.service.Me(
		r.Context(),
		userID,
	)

	if err != nil {
		response.Error(
			w,
			http.StatusNotFound,
			"user not found",
			nil,
		)
		return
	}

	response.Success(
		w,
		http.StatusOK,
		"user retrieved successfully",
		user,
	)
}
