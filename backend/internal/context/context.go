package context

import (
	stdcontext "context"

	"github.com/google/uuid"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func SetUserID(
	ctx stdcontext.Context,
	userID uuid.UUID,
) stdcontext.Context {
	return stdcontext.WithValue(
		ctx,
		UserIDKey,
		userID,
	)
}

func GetUserID(
	ctx stdcontext.Context,
) (uuid.UUID, bool) {
	userID, ok := ctx.Value(UserIDKey).(uuid.UUID)
	return userID, ok
}
