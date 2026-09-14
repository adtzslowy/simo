package middleware

import (
	"net/http"
	"strings"

	"github.com/adtzslowy/simo/internal/auth"
	appcontext "github.com/adtzslowy/simo/internal/context"
	"github.com/adtzslowy/simo/internal/response"
	"github.com/google/uuid"
)

type AuthMiddleware struct {
	jwtService *auth.JWTService
}

func NewAuthMiddleware(
	jwtService *auth.JWTService,
) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (m *AuthMiddleware) RequireAuth(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			response.Error(
				w, http.StatusUnauthorized, "authorization header is required", nil,
			)
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Error(
				w, http.StatusUnauthorized, "invalid authorization header", nil,
			)
			return
		}

		token := parts[1]

		claims, err := m.jwtService.ParseToken(token)
		if err != nil {
			response.Error(
				w, http.StatusUnauthorized, "invalid or expired token", nil,
			)
			return
		}

		userID := claims.UserID
		if userID == uuid.Nil {
			response.Error(
				w, http.StatusUnauthorized, "invalid user identity", nil,
			)
			return
		}

		ctx := appcontext.SetUserID(
			r.Context(),
			userID,
		)

		next.ServeHTTP(
			w, r.WithContext(ctx),
		)
	})
}
