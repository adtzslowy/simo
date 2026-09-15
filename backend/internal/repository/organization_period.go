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

func (r *OrganizationPeriodRepository) FindByID(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) (*model.OrganizationPeriod, error) {
	var period model.OrganizationPeriod

	err := r.db.QueryRow(ctx, `
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
		WHERE id = $1
		  AND organization_id = $2
	`, id, organizationID).Scan(
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

	return &period, nil
}

func (r *OrganizationPeriodRepository) Create(
	ctx context.Context,
	period *model.OrganizationPeriod,
) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO organization_periods (
			organization_id,
			name,
			start_date,
			end_date,
			is_active
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id,
			organization_id,
			name,
			start_date,
			end_date,
			is_active,
			created_at,
			updated_at
	`,
		period.OrganizationID,
		period.Name,
		period.StartDate,
		period.EndDate,
		period.IsActive,
	).Scan(
		&period.ID,
		&period.OrganizationID,
		&period.Name,
		&period.StartDate,
		&period.EndDate,
		&period.IsActive,
		&period.CreatedAt,
		&period.UpdatedAt,
	)
}

func (r *OrganizationPeriodRepository) Update(
	ctx context.Context,
	period *model.OrganizationPeriod,
) error {
	return r.db.QueryRow(ctx, `
		UPDATE organization_periods
		SET
			name = $1,
			start_date = $2,
			end_date = $3,
			is_active = $4,
			updated_at = NOW()
		WHERE id = $5
		  AND organization_id = $6
		RETURNING
			id,
			organization_id,
			name,
			start_date,
			end_date,
			is_active,
			created_at,
			updated_at
	`,
		period.Name,
		period.StartDate,
		period.EndDate,
		period.IsActive,
		period.ID,
		period.OrganizationID,
	).Scan(
		&period.ID,
		&period.OrganizationID,
		&period.Name,
		&period.StartDate,
		&period.EndDate,
		&period.IsActive,
		&period.CreatedAt,
		&period.UpdatedAt,
	)
}

func (r *OrganizationPeriodRepository) Delete(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM organization_periods
		WHERE id = $1
		  AND organization_id = $2
	`, id, organizationID)

	return err
}
