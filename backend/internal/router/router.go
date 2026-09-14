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
	rbacMiddleware *middleware.RBACMiddleware,
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

	mux.Handle(
		"GET /api/v1/organization-types",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("organization_type.read")(
				http.HandlerFunc(organizationTypeHandler.GetAll),
			),
		),
	)

	mux.Handle(
		"GET /api/v1/organizations",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("organization.read")(
				http.HandlerFunc(organizationHandler.GetAll),
			),
		),
	)

	mux.Handle(
		"GET /api/v1/organizations/{id}",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("organization.read")(
				http.HandlerFunc(organizationHandler.GetByID),
			),
		),
	)

	mux.Handle(
		"GET /api/v1/organizations/{organization_id}/periods",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("organization_period.read")(
				http.HandlerFunc(organizationPeriodHandler.GetByOrganizationID),
			),
		),
	)

	return mux
}
