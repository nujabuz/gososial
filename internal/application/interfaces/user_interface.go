// Package interfaces
package interfaces

import (
	"context"

	"github.com/nujabuz/gososial/internal/application/params"
	"github.com/nujabuz/gososial/internal/application/results"
)

type UserService interface {
	CreateUser(ctx context.Context, userParam *params.UserCreateParam) (*results.UserResult, error)
	FindUserByID(ctx context.Context, userParam *params.UserFindByIDParam) (*results.UserResult, error)
	FindUserByUsername(ctx context.Context, userParam *params.UserFindByUsernameParam) (*results.UserResult, error)
}
