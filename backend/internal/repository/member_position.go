package repository

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MemberPositionRepository struct {
	db *pgxpool.Pool
}

func NewMemberPositionRepository(
	db *pgxpool.Pool,
) *MemberPositionRepository {
	return &MemberPositionRepository{
		db: db,
	}
}

func (r *MemberPositionRepository) FindAllByMemberID(
	ctx context.Context,
	organizationID uuid.UUID,
	memberID uuid.UUID,
) ([]model.MemberPosition, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			mp.id,
			mp.member_id,
			mp.position_id,
			mp.department_id,
			mp.start_date,
			mp.end_date,
			mp.is_active,
			mp.created_at,
			mp.updated_at
		FROM member_positions mp
		JOIN organization_members om
			ON om.id = mp.member_id
		WHERE mp.member_id = $1
		  AND om.organization_id = $2
		ORDER BY mp.start_date DESC NULLS LAST, mp.created_at DESC
	`, memberID, organizationID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	assignments := make([]model.MemberPosition, 0)

	for rows.Next() {
		var assignment model.MemberPosition

		if err := rows.Scan(
			&assignment.ID,
			&assignment.MemberID,
			&assignment.PositionID,
			&assignment.DepartmentID,
			&assignment.StartDate,
			&assignment.EndDate,
			&assignment.IsActive,
			&assignment.CreatedAt,
			&assignment.UpdatedAt,
		); err != nil {
			return nil, err
		}

		assignments = append(assignments, assignment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return assignments, nil
}

func (r *MemberPositionRepository) FindByID(
	ctx context.Context,
	organizationID uuid.UUID,
	memberID uuid.UUID,
	id uuid.UUID,
) (*model.MemberPosition, error) {
	var assignment model.MemberPosition

	err := r.db.QueryRow(ctx, `
		SELECT
			mp.id,
			mp.member_id,
			mp.position_id,
			mp.department_id,
			mp.start_date,
			mp.end_date,
			mp.is_active,
			mp.created_at,
			mp.updated_at
		FROM member_positions mp
		JOIN organization_members om
			ON om.id = mp.member_id
		WHERE mp.id = $1
		  AND mp.member_id = $2
		  AND om.organization_id = $3
	`, id, memberID, organizationID).Scan(
		&assignment.ID,
		&assignment.MemberID,
		&assignment.PositionID,
		&assignment.DepartmentID,
		&assignment.StartDate,
		&assignment.EndDate,
		&assignment.IsActive,
		&assignment.CreatedAt,
		&assignment.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &assignment, nil
}

func (r *MemberPositionRepository) Create(
	ctx context.Context,
	organizationID uuid.UUID,
	memberID uuid.UUID,
	assignment *model.MemberPosition,
) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO member_positions (
			member_id,
			position_id,
			department_id,
			start_date,
			end_date,
			is_active
		)
		SELECT
			om.id,
			p.id,
			d.id,
			$4,
			$5,
			$6
		FROM organization_members om
		JOIN positions p
			ON p.id = $3
		   AND p.organization_id = om.organization_id
		LEFT JOIN departments d
			ON d.id = $7
		   AND d.organization_id = om.organization_id
		WHERE om.id = $2
		  AND om.organization_id = $1
		RETURNING
			id,
			member_id,
			position_id,
			department_id,
			start_date,
			end_date,
			is_active,
			created_at,
			updated_at
	`,
		organizationID,
		memberID,
		assignment.PositionID,
		assignment.StartDate,
		assignment.EndDate,
		assignment.IsActive,
		assignment.DepartmentID,
	).Scan(
		&assignment.ID,
		&assignment.MemberID,
		&assignment.PositionID,
		&assignment.DepartmentID,
		&assignment.StartDate,
		&assignment.EndDate,
		&assignment.IsActive,
		&assignment.CreatedAt,
		&assignment.UpdatedAt,
	)
}

func (r *MemberPositionRepository) Update(
	ctx context.Context,
	organizationID uuid.UUID,
	memberID uuid.UUID,
	id uuid.UUID,
	assignment *model.MemberPosition,
) error {
	return r.db.QueryRow(ctx, `
		UPDATE member_positions mp
		SET
			position_id = p.id,
			department_id = d.id,
			start_date = $5,
			end_date = $6,
			is_active = $7,
			updated_at = NOW()
		FROM organization_members om
		JOIN positions p
			ON p.id = $4
		   AND p.organization_id = om.organization_id
		LEFT JOIN departments d
			ON d.id = $3
		   AND d.organization_id = om.organization_id
		WHERE mp.id = $2
		  AND mp.member_id = $1
		  AND om.id = mp.member_id
		  AND om.organization_id = $8
		RETURNING
			mp.id,
			mp.member_id,
			mp.position_id,
			mp.department_id,
			mp.start_date,
			mp.end_date,
			mp.is_active,
			mp.created_at,
			mp.updated_at
	`,
		memberID,
		id,
		assignment.DepartmentID,
		assignment.PositionID,
		assignment.StartDate,
		assignment.EndDate,
		assignment.IsActive,
		organizationID,
	).Scan(
		&assignment.ID,
		&assignment.MemberID,
		&assignment.PositionID,
		&assignment.DepartmentID,
		&assignment.StartDate,
		&assignment.EndDate,
		&assignment.IsActive,
		&assignment.CreatedAt,
		&assignment.UpdatedAt,
	)
}

func (r *MemberPositionRepository) Delete(
	ctx context.Context,
	organizationID uuid.UUID,
	memberID uuid.UUID,
	id uuid.UUID,
) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM member_positions mp
		USING organization_members om
		WHERE mp.id = $1
		  AND mp.member_id = $2
		  AND om.id = mp.member_id
		  AND om.organization_id = $3
	`, id, memberID, organizationID)

	return err
}
