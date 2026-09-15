package repository

import (
	"context"

	"github.com/adtzslowy/simo/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganizationMemberRepository struct {
	db *pgxpool.Pool
}

func NewOrganizationMemberRepository(
	db *pgxpool.Pool,
) *OrganizationMemberRepository {
	return &OrganizationMemberRepository{
		db: db,
	}
}

func (r *OrganizationMemberRepository) FindAllByOrganizationID(
	ctx context.Context,
	organizationID uuid.UUID,
) ([]model.OrganizationMember, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			organization_id,
			user_id,
			period_id,
			status,
			joined_at,
			left_at,
			created_at,
			updated_at
		FROM organization_members
		WHERE organization_id = $1
		ORDER BY created_at DESC
	`, organizationID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	members := make([]model.OrganizationMember, 0)

	for rows.Next() {
		var member model.OrganizationMember

		err := rows.Scan(
			&member.ID,
			&member.OrganizationID,
			&member.UserID,
			&member.PeriodID,
			&member.Status,
			&member.JoinedAt,
			&member.LeftAt,
			&member.CreatedAt,
			&member.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func (r *OrganizationMemberRepository) FindByID(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) (*model.OrganizationMember, error) {
	var member model.OrganizationMember

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			organization_id,
			user_id,
			period_id,
			status,
			joined_at,
			left_at,
			created_at,
			updated_at
		FROM organization_members
		WHERE id = $1
		  AND organization_id = $2
	`, id, organizationID).Scan(
		&member.ID,
		&member.OrganizationID,
		&member.UserID,
		&member.PeriodID,
		&member.Status,
		&member.JoinedAt,
		&member.LeftAt,
		&member.CreatedAt,
		&member.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &member, nil
}

func (r *OrganizationMemberRepository) Create(
	ctx context.Context,
	member *model.OrganizationMember,
) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO organization_members (
			organization_id,
			user_id,
			period_id,
			status,
			joined_at,
			left_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING
			id,
			organization_id,
			user_id,
			period_id,
			status,
			joined_at,
			left_at,
			created_at,
			updated_at
	`,
		member.OrganizationID,
		member.UserID,
		member.PeriodID,
		member.Status,
		member.JoinedAt,
		member.LeftAt,
	).Scan(
		&member.ID,
		&member.OrganizationID,
		&member.UserID,
		&member.PeriodID,
		&member.Status,
		&member.JoinedAt,
		&member.LeftAt,
		&member.CreatedAt,
		&member.UpdatedAt,
	)
}

func (r *OrganizationMemberRepository) Update(
	ctx context.Context,
	member *model.OrganizationMember,
) error {
	return r.db.QueryRow(ctx, `
		UPDATE organization_members
		SET
			period_id = $1,
			status = $2,
			joined_at = $3,
			left_at = $4,
			updated_at = NOW()
		WHERE id = $5
		  AND organization_id = $6
		RETURNING
			id,
			organization_id,
			user_id,
			period_id,
			status,
			joined_at,
			left_at,
			created_at,
			updated_at
	`,
		member.PeriodID,
		member.Status,
		member.JoinedAt,
		member.LeftAt,
		member.ID,
		member.OrganizationID,
	).Scan(
		&member.ID,
		&member.OrganizationID,
		&member.UserID,
		&member.PeriodID,
		&member.Status,
		&member.JoinedAt,
		&member.LeftAt,
		&member.CreatedAt,
		&member.UpdatedAt,
	)
}

func (r *OrganizationMemberRepository) Delete(
	ctx context.Context,
	organizationID uuid.UUID,
	id uuid.UUID,
) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM organization_members
		WHERE id = $1
		  AND organization_id = $2
	`, id, organizationID)

	return err
}
