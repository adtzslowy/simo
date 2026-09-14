package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RBACRepository struct {
	db *pgxpool.Pool
}

func NewRBACRepository(db *pgxpool.Pool) *RBACRepository {
	return &RBACRepository{
		db: db,
	}
}

func (r *RBACRepository) HasPermission(
	ctx context.Context,
	userID uuid.UUID,
	permissionCode string,
) (bool, error) {
	var exists bool

	err := r.db.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1
			FROM user_roles ur
			JOIN roles r
				ON r.id = ur.role_id
			JOIN role_permissions rp
				ON rp.role_id = r.id
			JOIN permissions p
				ON p.id = rp.permission_id
			WHERE ur.user_id = $1
			  AND p.code = $2
			  AND r.is_active = TRUE
			  AND p.is_active = TRUE
		)
	`, userID, permissionCode).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
