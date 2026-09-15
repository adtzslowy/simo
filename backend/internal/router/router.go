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
	memberPositionHandler *handler.MemberPositionHandler,
	authHandler *handler.AuthHandler,
	authMiddleware *middleware.AuthMiddleware,
) http.Handler {
	mux := http.NewServeMux()

	// =========================================================
	// AUTH
	// =========================================================

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

	// =========================================================
	// ORGANIZATION TYPES
	// =========================================================

	mux.Handle(
		"GET /api/v1/organization-types",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationTypeHandler.GetAll),
		),
	)

	mux.Handle(
		"GET /api/v1/organization-types/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationTypeHandler.GetByID),
		),
	)

	mux.Handle(
		"POST /api/v1/organization-types",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationTypeHandler.Create),
		),
	)

	mux.Handle(
		"PUT /api/v1/organization-types/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationTypeHandler.Update),
		),
	)

	mux.Handle(
		"DELETE /api/v1/organization-types/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationTypeHandler.Delete),
		),
	)

	// =========================================================
	// ORGANIZATIONS
	// =========================================================

	mux.Handle(
		"GET /api/v1/organizations",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationHandler.GetAll),
		),
	)

	mux.Handle(
		"GET /api/v1/organizations/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationHandler.GetByID),
		),
	)

	mux.Handle(
		"POST /api/v1/organizations",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationHandler.Create),
		),
	)

	mux.Handle(
		"PUT /api/v1/organizations/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationHandler.Update),
		),
	)

	mux.Handle(
		"DELETE /api/v1/organizations/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationHandler.Delete),
		),
	)

	// =========================================================
	// ORGANIZATION PERIODS
	// =========================================================

	mux.Handle(
		"GET /api/v1/organizations/{organization_id}/periods",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationPeriodHandler.GetByOrganizationID),
		),
	)

	mux.Handle(
		"GET /api/v1/organizations/{organization_id}/periods/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationPeriodHandler.GetByID),
		),
	)

	mux.Handle(
		"POST /api/v1/organizations/{organization_id}/periods",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationPeriodHandler.Create),
		),
	)

	mux.Handle(
		"PUT /api/v1/organizations/{organization_id}/periods/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationPeriodHandler.Update),
		),
	)

	mux.Handle(
		"DELETE /api/v1/organizations/{organization_id}/periods/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(organizationPeriodHandler.Delete),
		),
	)

	// =========================================================
	// MEMBER POSITION ASSIGNMENTS
	// =========================================================

	mux.Handle(
		"GET /api/v1/organizations/{organization_id}/members/{member_id}/positions",
		authMiddleware.RequireAuth(
			http.HandlerFunc(memberPositionHandler.GetAll),
		),
	)

	mux.Handle(
		"GET /api/v1/organizations/{organization_id}/members/{member_id}/positions/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(memberPositionHandler.GetByID),
		),
	)

	mux.Handle(
		"POST /api/v1/organizations/{organization_id}/members/{member_id}/positions",
		authMiddleware.RequireAuth(
			http.HandlerFunc(memberPositionHandler.Create),
		),
	)

	mux.Handle(
		"PUT /api/v1/organizations/{organization_id}/members/{member_id}/positions/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(memberPositionHandler.Update),
		),
	)

	mux.Handle(
		"DELETE /api/v1/organizations/{organization_id}/members/{member_id}/positions/{id}",
		authMiddleware.RequireAuth(
			http.HandlerFunc(memberPositionHandler.Delete),
		),
	)

	return mux
}
