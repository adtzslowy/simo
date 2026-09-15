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
	departmentHandler *handler.DepartmentHandler,
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
		"POST /api/v1/organization-types",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("organization_type.create")(
				http.HandlerFunc(organizationTypeHandler.Create),
			),
		),
	)

	mux.Handle(
		"PUT /api/v1/organization-types/{id}",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("organization_type.update")(
				http.HandlerFunc(organizationTypeHandler.Update),
			),
		),
	)

	mux.Handle(
		"DELETE /api/v1/organization-types/{id}/delete",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("organization_type.delete")(
				http.HandlerFunc(organizationTypeHandler.Delete),
			),
		),
	)

	mux.Handle(
		"POST /api/v1/organizations",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("organization.create")(
				http.HandlerFunc(organizationHandler.Create),
			),
		),
	)

	mux.Handle(
		"PUT /api/v1/organizations/{id}",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("organization.update")(
				http.HandlerFunc(organizationHandler.Update),
			),
		),
	)

	mux.Handle(
		"DELETE /api/v1/organizations/{id}",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("organization.delete")(
				http.HandlerFunc(organizationHandler.Delete),
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
			rbacMiddleware.RequirePermission("period.read")(
				http.HandlerFunc(organizationPeriodHandler.GetByOrganizationID),
			),
		),
	)

	mux.Handle(
		"GET /api/v1/organizations/{organization_id}/periods/{id}",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("period.read")(
				http.HandlerFunc(organizationPeriodHandler.GetByID),
			),
		),
	)

	mux.Handle(
		"POST /api/v1/organizations/{organization_id}/periods",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("period.create")(
				http.HandlerFunc(organizationPeriodHandler.Create),
			),
		),
	)

	mux.Handle(
		"PUT /api/v1/organizations/{organization_id}/periods/{id}",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("period.update")(
				http.HandlerFunc(organizationPeriodHandler.Update),
			),
		),
	)

	mux.Handle(
		"DELETE /api/v1/organizations/{organization_id}/periods/{id}",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("period.delete")(
				http.HandlerFunc(organizationPeriodHandler.Delete),
			),
		),
	)

	mux.Handle(
		"GET /api/v1/organizations/{organization_id}/departments",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("department.read")(
				http.HandlerFunc(
					departmentHandler.GetAllByOrganizationID,
				),
			),
		),
	)

	mux.Handle(
		"GET /api/v1/organizations/{organization_id}/departments/{id}",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("department.read")(
				http.HandlerFunc(
					departmentHandler.GetByID,
				),
			),
		),
	)

	mux.Handle(
		"POST /api/v1/organizations/{organization_id}/departments",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("department.create")(
				http.HandlerFunc(
					departmentHandler.Create,
				),
			),
		),
	)

	mux.Handle(
		"PUT /api/v1/organizations/{organization_id}/departments/{id}",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("department.update")(
				http.HandlerFunc(
					departmentHandler.Update,
				),
			),
		),
	)

	mux.Handle(
		"DELETE /api/v1/organizations/{organization_id}/departments/{id}",
		authMiddleware.RequireAuth(
			rbacMiddleware.RequirePermission("department.delete")(
				http.HandlerFunc(
					departmentHandler.Delete,
				),
			),
		),
	)

	return mux
}
