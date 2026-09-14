package repository

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganizationPeriodRepository struct {
	db *pgxpool.Pool
}

func NewOrganizationPeriodRepository(db *pgxpool.Pool) *OrganizationPeriodRepository {
	return &OrganizationPeriodRepository{
		db: db,
	}
}

func (r *OrganizationPeriodRepository) FindByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]model.OrganizationPeriod, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			organization_id,
			name,
			start_date,
			end_date,
			is_active,
			created_at,
			updated_at
		FROM organization_periods
		WHERE organization_id = $1
		ORDER BY start_date DESC
	`, organizationID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	periods := make([]model.OrganizationPeriod, 0)

	for rows.Next() {
		var period model.OrganizationPeriod

		err := rows.Scan(
			&period.ID,
			&period.OrganizationID,
			&period.Name,
			&period.StartDate,
			&period.EndDate,
			&period.IsActive,
			&period.CreatedAt,
			&period.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		periods = append(periods, period)

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return periods, nil
}
