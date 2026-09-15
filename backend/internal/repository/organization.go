package repository

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganizationRepository struct {
	db *pgxpool.Pool
}

func NewOrganizationRepository(db *pgxpool.Pool) *OrganizationRepository {
	return &OrganizationRepository{
		db: db,
	}
}

func (r *OrganizationRepository) FindAll(ctx context.Context) ([]model.Organization, error) {
	rows, err := r.db.Query(ctx,
		`SELECT
			id,
			organization_type_id,
			institution_id,
			academic_department_id,
			name,
			code,
			description,
			logo_url,
			is_active,
			created_at,
			updated_at
		FROM organizations
		ORDER BY name ASC`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var organizations []model.Organization

	for rows.Next() {
		var organization model.Organization

		err := rows.Scan(
			&organization.ID,
			&organization.OrganizationTypeID,
			&organization.InstitutionID,
			&organization.AcademicDepartmentID,
			&organization.Name,
			&organization.Code,
			&organization.Description,
			&organization.LogoURL,
			&organization.IsActive,
			&organization.CreatedAt,
			&organization.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		organizations = append(organizations, organization)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return organizations, nil
}

func (r *OrganizationRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.Organization, error) {
	var organization model.Organization

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			organization_type_id,
			institution_id,
			academic_department_id,
			name,
			code,
			description,
			logo_url,
			is_active,
			created_at,
			updated_at
		FROM organizations
		WHERE id = $1
	`, id).Scan(
		&organization.ID,
		&organization.OrganizationTypeID,
		&organization.InstitutionID,
		&organization.AcademicDepartmentID,
		&organization.Name,
		&organization.Code,
		&organization.Description,
		&organization.LogoURL,
		&organization.IsActive,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}

func (r *OrganizationRepository) Create(
	ctx context.Context,
	organization *model.Organization,
) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO organizations (
			organization_type_id,
			institution_id,
			academic_department_id,
			name,
			code,
			description,
			logo_url,
			is_active
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
		RETURNING
			id,
			organization_type_id,
			institution_id,
			academic_department_id,
			name,
			code,
			description,
			logo_url,
			is_active,
			created_at,
			updated_at
	`,
		organization.OrganizationTypeID,
		organization.InstitutionID,
		organization.AcademicDepartmentID,
		organization.Name,
		organization.Code,
		organization.Description,
		organization.LogoURL,
		organization.IsActive,
	).Scan(
		&organization.ID,
		&organization.OrganizationTypeID,
		&organization.InstitutionID,
		&organization.AcademicDepartmentID,
		&organization.Name,
		&organization.Code,
		&organization.Description,
		&organization.LogoURL,
		&organization.IsActive,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)
}

func (r *OrganizationRepository) Update(
	ctx context.Context,
	organization *model.Organization,
) error {
	return r.db.QueryRow(ctx, `
		UPDATE organizations
		SET
			organization_type_id = $1,
			institution_id = $2,
			academic_department_id = $3,
			name = $4,
			code = $5,
			description = $6,
			logo_url = $7,
			is_active = $8,
			updated_at = NOW()
		WHERE id = $9
		RETURNING
			id,
			organization_type_id,
			institution_id,
			academic_department_id,
			name,
			code,
			description,
			logo_url,
			is_active,
			created_at,
			updated_at
	`,
		organization.OrganizationTypeID,
		organization.InstitutionID,
		organization.AcademicDepartmentID,
		organization.Name,
		organization.Code,
		organization.Description,
		organization.LogoURL,
		organization.IsActive,
		organization.ID,
	).Scan(
		&organization.ID,
		&organization.OrganizationTypeID,
		&organization.InstitutionID,
		&organization.AcademicDepartmentID,
		&organization.Name,
		&organization.Code,
		&organization.Description,
		&organization.LogoURL,
		&organization.IsActive,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)
}

func (r *OrganizationRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM organizations
		WHERE id = $1
	`, id)

	return err
}
