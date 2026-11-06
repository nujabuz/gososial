// Package usecase business layer
package usecase

import (
	"context"

	"github.com/nujabuz/gososial/internal/entity"
)

type UserUseCase interface {
	Delete(c context.Context, in *InputDeleteUser) error
	Create(c context.Context, in *InputCreateUser) (*OutputUser, error)
	Update(c context.Context, in *InputUpdateUser) (*OutputUser, error)
	FindByID(c context.Context, in *InputFindByIDUser) (*OutputUser, error)
	FindByUsername(c context.Context, in *InputFindByUsernameUser) (*OutputUser, error)
}

type PostUseCase interface {
	Delete(c context.Context, in *InputDeletePost) error
	Create(c context.Context, in *InputCreatePost) (*OutputPost, error)
	Update(c context.Context, in *InputUpdatePost) (*OutputPost, error)
	FindByID(c context.Context, in *InputFindByIDPost) (*OutputPost, error)
	FindByUserID(c context.Context, in *InputFindByUserIDPost) (*OutputPost, error)
	FindManyByUserID(c context.Context, in *InputFindByUserIDPost) ([]*OutputPost, error)
}

type UserRepository interface {
	Delete(c context.Context, id int64) error
	Create(c context.Context, user *entity.User) error
	Update(c context.Context, user *entity.User) error
	FindByID(c context.Context, id int64) (*entity.User, error)
	FindByUsername(c context.Context, username string) (*entity.User, error)
}

type PostRepository interface {
	Delete(c context.Context, id int64, userID int64) error
	Create(c context.Context, post *entity.Post) error
	Update(c context.Context, post *entity.Post) error
	FindByID(c context.Context, id int64) (*entity.Post, error)
	FindByUserID(c context.Context, userID int64) (*entity.Post, error)
	FindManyByUserID(c context.Context, userID int64) ([]*entity.Post, error)
}

type HasherPassword interface {
	HashPassword(plainPassword string) (string, error)
	CompareHashedPassword(plainPassword string, hashedPassword string) error
}
