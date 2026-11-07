// Package services
package services

import (
	"context"

	"github.com/nujabuz/gososial/internal/application/interfaces"
	"github.com/nujabuz/gososial/internal/application/params"
	"github.com/nujabuz/gososial/internal/application/results"
	"github.com/nujabuz/gososial/internal/domain/entities"
	"github.com/nujabuz/gososial/internal/domain/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) interfaces.UserService {
	return &userService{
		repo: userRepository,
	}
}

func (service *userService) CreateUser(
	ctx context.Context,
	userParam *params.UserCreateParam,
) (*results.UserResult, error) {
	user := entities.NewUser(
		userParam.Username,
		userParam.Password,
	)
	if err := user.Validate(); err != nil {
		return nil, err
	}

	createdUser, err := service.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return results.UserToResult(createdUser), nil
}

func (service *userService) FindUserByID(
	ctx context.Context,
	userParam *params.UserFindByIDParam,
) (*results.UserResult, error) {
	user, err := service.repo.FindByID(ctx, userParam.ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entities.ErrNotFound
	}

	return results.UserToResult(user), nil
}

func (service *userService) FindUserByUsername(
	ctx context.Context,
	userParam *params.UserFindByUsernameParam,
) (*results.UserResult, error) {
	user, err := service.repo.FindByUsername(ctx, userParam.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entities.ErrNotFound
	}

	return results.UserToResult(user), nil
}
