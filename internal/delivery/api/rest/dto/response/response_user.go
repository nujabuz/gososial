// Package response
package response

import (
	"time"

	"github.com/nujabuz/gososial/internal/application/results"
)

type ResponseUser struct {
	ID        int64
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewResponseUser(user *results.UserResult) ResponseUser {
	return ResponseUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
