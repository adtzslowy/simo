package repository

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DepartmentRepository struct {
	db *pgxpool.Pool
}

func NewDepartmentRepository(
	db *pgxpool.Pool,
) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}

func (r *DepartmentRepository) FindAllByOrganizationID(
	ctx context.Context,
	organizationID uuid.UUID,
) ([]model.Department, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			organization_id,
			period_id,
			parent_id,
			name,
			code,
			description,
			is_active,
			created_at,
			updated_at
		FROM departments
		WHERE organization_id = $1
		ORDER BY name ASC
	`, organizationID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	departments := make([]model.Department, 0)

	for rows.Next() {
		var department model.Department

		err := rows.Scan(
			&department.ID,
			&department.OrganizationID,
			&department.PeriodID,
			&department.ParentID,
			&department.Name,
			&department.Code,
			&department.Description,
			&department.IsActive,
			&department.CreatedAt,
			&department.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		departments = append(departments, department)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *DepartmentRepository) FindByID(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) (*model.Department, error) {
	var department model.Department

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			organization_id,
			period_id,
			parent_id,
			name,
			code,
			description,
			is_active,
			created_at,
			updated_at
		FROM departments
		WHERE id = $1
		  AND organization_id = $2
	`, id, organizationID).Scan(
		&department.ID,
		&department.OrganizationID,
		&department.PeriodID,
		&department.ParentID,
		&department.Name,
		&department.Code,
		&department.Description,
		&department.IsActive,
		&department.CreatedAt,
		&department.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *DepartmentRepository) Create(
	ctx context.Context,
	department *model.Department,
) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO departments (
			organization_id,
			period_id,
			parent_id,
			name,
			code,
			description,
			is_active
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			id,
			organization_id,
			period_id,
			parent_id,
			name,
			code,
			description,
			is_active,
			created_at,
			updated_at
	`,
		department.OrganizationID,
		department.PeriodID,
		department.ParentID,
		department.Name,
		department.Code,
		department.Description,
		department.IsActive,
	).Scan(
		&department.ID,
		&department.OrganizationID,
		&department.PeriodID,
		&department.ParentID,
		&department.Name,
		&department.Code,
		&department.Description,
		&department.IsActive,
		&department.CreatedAt,
		&department.UpdatedAt,
	)
}

func (r *DepartmentRepository) Update(
	ctx context.Context,
	department *model.Department,
) error {
	return r.db.QueryRow(ctx, `
		UPDATE departments
		SET
			period_id = $1,
			parent_id = $2,
			name = $3,
			code = $4,
			description = $5,
			is_active = $6,
			updated_at = NOW()
		WHERE id = $7
		  AND organization_id = $8
		RETURNING
			id,
			organization_id,
			period_id,
			parent_id,
			name,
			code,
			description,
			is_active,
			created_at,
			updated_at
	`,
		department.PeriodID,
		department.ParentID,
		department.Name,
		department.Code,
		department.Description,
		department.IsActive,
		department.ID,
		department.OrganizationID,
	).Scan(
		&department.ID,
		&department.OrganizationID,
		&department.PeriodID,
		&department.ParentID,
		&department.Name,
		&department.Code,
		&department.Description,
		&department.IsActive,
		&department.CreatedAt,
		&department.UpdatedAt,
	)
}

func (r *DepartmentRepository) Delete(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM departments
		WHERE id = $1
		  AND organization_id = $2
	`, id, organizationID)

	return err
}
