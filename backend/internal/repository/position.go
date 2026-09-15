package repository

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PositionRepository struct {
	db *pgxpool.Pool
}

func NewPositionRepository(
	db *pgxpool.Pool,
) *PositionRepository {
	return &PositionRepository{
		db: db,
	}
}

func (r *PositionRepository) FindAllByOrganizationID(
	ctx context.Context,
	organizationID uuid.UUID,
) ([]model.Position, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			organization_id,
			name,
			code,
			description,
			is_active,
			created_at,
			updated_at
		FROM positions
		WHERE organization_id = $1
		ORDER BY name ASC
	`, organizationID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	positions := make([]model.Position, 0)

	for rows.Next() {
		var position model.Position

		err := rows.Scan(
			&position.ID,
			&position.OrganizationID,
			&position.Name,
			&position.Code,
			&position.Description,
			&position.IsActive,
			&position.CreatedAt,
			&position.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		positions = append(positions, position)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}

func (r *PositionRepository) FindByID(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) (*model.Position, error) {
	var position model.Position

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			organization_id,
			name,
			code,
			description,
			is_active,
			created_at,
			updated_at
		FROM positions
		WHERE id = $1
		  AND organization_id = $2
	`, id, organizationID).Scan(
		&position.ID,
		&position.OrganizationID,
		&position.Name,
		&position.Code,
		&position.Description,
		&position.IsActive,
		&position.CreatedAt,
		&position.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &position, nil
}

func (r *PositionRepository) Create(
	ctx context.Context,
	position *model.Position,
) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO positions (
			organization_id,
			name,
			code,
			description,
			is_active
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id,
			organization_id,
			name,
			code,
			description,
			is_active,
			created_at,
			updated_at
	`,
		position.OrganizationID,
		position.Name,
		position.Code,
		position.Description,
		position.IsActive,
	).Scan(
		&position.ID,
		&position.OrganizationID,
		&position.Name,
		&position.Code,
		&position.Description,
		&position.IsActive,
		&position.CreatedAt,
		&position.UpdatedAt,
	)
}

func (r *PositionRepository) Update(
	ctx context.Context,
	position *model.Position,
) error {
	return r.db.QueryRow(ctx, `
		UPDATE positions
		SET
			name = $1,
			code = $2,
			description = $3,
			is_active = $4,
			updated_at = NOW()
		WHERE id = $5
		  AND organization_id = $6
		RETURNING
			id,
			organization_id,
			name,
			code,
			description,
			is_active,
			created_at,
			updated_at
	`,
		position.Name,
		position.Code,
		position.Description,
		position.IsActive,
		position.ID,
		position.OrganizationID,
	).Scan(
		&position.ID,
		&position.OrganizationID,
		&position.Name,
		&position.Code,
		&position.Description,
		&position.IsActive,
		&position.CreatedAt,
		&position.UpdatedAt,
	)
}

func (r *PositionRepository) Delete(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM positions
		WHERE id = $1
		  AND organization_id = $2
	`, id, organizationID)

	return err
}
