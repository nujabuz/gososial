// Package results
package results

import (
	"time"

	"github.com/nujabuz/gososial/internal/domain/entities"
)

type UserResult struct {
	ID        int64
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserToResult(user *entities.User) *UserResult {
	return &UserResult{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
