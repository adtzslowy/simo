package repository

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganizationTypeRepository struct {
	db *pgxpool.Pool
}

func NewOrganizationTypeRepository(db *pgxpool.Pool) *OrganizationTypeRepository {
	return &OrganizationTypeRepository{
		db: db,
	}
}

func (r *OrganizationTypeRepository) FindAll(
	ctx context.Context,
) ([]model.OrganizationType, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			name,
			code,
			description,
			created_at,
			updated_at
		FROM organization_types
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []model.OrganizationType

	for rows.Next() {
		var organizationType model.OrganizationType

		err := rows.Scan(
			&organizationType.ID,
			&organizationType.Name,
			&organizationType.Code,
			&organizationType.Description,
			&organizationType.CreatedAt,
			&organizationType.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		types = append(types, organizationType)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return types, nil
}

func (r *OrganizationTypeRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.OrganizationType, error) {
	var organizationType model.OrganizationType

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			name,
			code,
			description,
			created_at,
			updated_at
		FROM organization_types
		WHERE id = $1
	`, id).Scan(
		&organizationType.ID,
		&organizationType.Name,
		&organizationType.Code,
		&organizationType.Description,
		&organizationType.CreatedAt,
		&organizationType.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &organizationType, nil
}

func (r *OrganizationTypeRepository) Create(ctx context.Context, organizationType *model.OrganizationType) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO organization_types (
			name,
			code,
			description
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			name,
			code,
			description,
			created_at,
			updated_at
	`, organizationType.Name, organizationType.Code, organizationType.Description).Scan(
		&organizationType.ID,
		&organizationType.Name,
		&organizationType.Code,
		&organizationType.Description,
		&organizationType.CreatedAt,
		&organizationType.UpdatedAt,
	)
}

func (r *OrganizationTypeRepository) Update(
	ctx context.Context,
	organizationType *model.OrganizationType,
) error {
	return r.db.QueryRow(ctx, `
		UPDATE organization_types
		SET
			name = $1,
			code = $2,
			description = $3,
			updated_at = NOW()
		WHERE id = $4
		RETURNING
			id,
			name,
			code,
			description,
			created_at,
			updated_at
	`, organizationType.Name, organizationType.Code, organizationType.Description, organizationType.ID).Scan(
		&organizationType.ID,
		&organizationType.Name,
		&organizationType.Code,
		&organizationType.Description,
		&organizationType.CreatedAt,
		&organizationType.UpdatedAt,
	)
}

func (r *OrganizationTypeRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM organization_types WHERE id = $1
	`, id)

	return err
}
