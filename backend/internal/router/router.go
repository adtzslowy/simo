package router

import (
	"net/http"

	"github.com/adtzslowy/simo/internal/handler"
)

func New(
	organizationTypeHandler *handler.OrganizationTypeHandler,
	organizationHandler *handler.OrganizationHandler,
	organizationPeriodHandler *handler.OrganizationPeriodHandler,
) http.Handler {
	mux := http.NewServeMux()

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
