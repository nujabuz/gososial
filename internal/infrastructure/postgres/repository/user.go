package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/nujabuz/gososial/internal/domain/entities"
	"github.com/nujabuz/gososial/internal/domain/repository"
	"github.com/nujabuz/gososial/internal/infrastructure/postgres/sqlc"
)

type userRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) repository.UserRepository {
	return &userRepository{
		queries: queries,
	}
}

func (repo *userRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	err := repo.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: timeToTimestamptz(user.CreatedAt),
		UpdatedAt: timeToTimestamptz(user.UpdatedAt),
	})
	if err != nil {
		return nil, nil
	}

	return nil, nil
}

func (repo *userRepository) FindByID(ctx context.Context, id int64) (*entities.User, error) {
	user, err := repo.queries.FindUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return modelToDomain(user), nil
}

func (repo *userRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error)  {
	user, err := repo.queries.FindUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return modelToDomain(user), nil
}

func modelToDomain(user sqlc.User) *entities.User {
	return &entities.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: timestamptzToTime(user.CreatedAt),
		UpdatedAt: timestamptzToTime(user.UpdatedAt),
	}
}
