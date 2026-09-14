package router

import (
	"net/http"

	"github.com/adtzslowy/simo/internal/handler"
	"github.com/adtzslowy/simo/internal/middleware"
)

func New(
	organizationTypeHandler *handler.OrganizationTypeHandler,
	organizationHandler *handler.OrganizationHandler,
	organizationPeriodHandler *handler.OrganizationPeriodHandler,
	authHandler *handler.AuthHandler,
	authMiddleware *middleware.AuthMiddleware,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"POST /api/v1/auth/login",
		authHandler.Login,
	)

	mux.Handle(
		"GET /api/v1/auth/me",
		authMiddleware.RequireAuth(
			http.HandlerFunc(authHandler.Me),
		),
	)

	mux.HandleFunc(
		"GET /api/v1/organization-types",
		organizationTypeHandler.GetAll,
	)

	mux.HandleFunc(
		"GET /api/v1/organizations",
		organizationHandler.GetAll,
	)

	mux.HandleFunc(
		"GET /api/v1/organizations/{id}",
		organizationHandler.GetByID,
	)

	mux.HandleFunc(
		"GET /api/v1/organizations/{organization_id}/periods",
		organizationPeriodHandler.GetByOrganizationID,
	)

	return mux
}
