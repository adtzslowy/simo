package middleware

import (
	"net/http"

	appcontext "github.com/adtzslowy/simo/internal/context"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/adtzslowy/simo/internal/response"
)

type RBACMiddleware struct {
	repository *repository.RBACRepository
}

func NewRBACMiddleware(repository *repository.RBACRepository) *RBACMiddleware {
	return &RBACMiddleware{
		repository: repository,
	}
}

func (m *RBACMiddleware) RequirePermission(permission string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, ok := appcontext.GetUserID(r.Context())
			if !ok {
				response.Error(w, http.StatusUnauthorized, "unauthorized", nil)
				return
			}

			hasPermission, err := m.repository.HasPermission(
				r.Context(),
				userID,
				permission,
			)

			if err != nil {
				response.Error(
					w, http.StatusInternalServerError, "failed to check permission", nil,
				)
				return
			}

			if !hasPermission {
				response.Error(
					w, http.StatusForbidden, "forbidden", nil,
				)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
