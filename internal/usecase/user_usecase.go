package usecase

import (
	"context"
	"errors"

	"github.com/nujabuz/gososial/internal/entity"
)

type userUseCase struct {
	// utility
	hasherPassword HasherPassword

	// repository
	userRepository UserRepository
}

func NewUserUseCase(hasherPassword HasherPassword, userRepository UserRepository) UserUseCase {
	return &userUseCase{
		hasherPassword: hasherPassword,
		userRepository: userRepository,
	}
}

func (uc *userUseCase) Delete(c context.Context, in *InputDeleteUser) error {
	// check user
	user, err := uc.userRepository.FindByID(c, in.ID)
	if err != nil {
		return err
	}

	// delete user
	err = uc.userRepository.Delete(c, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) Create(c context.Context, in *InputCreateUser) (*OutputUser, error) {
	// check user by username
	_, err := uc.userRepository.FindByUsername(c, in.Username)
	if err != nil && !errors.Is(err, entity.ErrNotFound) {
		return nil, err
	}

	// make new user
	user, err := entity.NewUser(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	// hash the password
	hashedPassword, err := uc.hasherPassword.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	// update user password
	if err = user.UpdatePassword(hashedPassword); err != nil {
		return nil, err
	}

	// create user
	if err = uc.userRepository.Create(c, user); err != nil {
		return nil, err
	}

	// change User to OutputUser
	out := uc.userToUserOutput(user)

	return out, nil
}

func (uc *userUseCase) Update(c context.Context, in *InputUpdateUser) (*OutputUser, error) {
	// check user
	user, err := uc.userRepository.FindByID(c, in.ID)
	if err != nil {
		return nil, err
	}

	// check value for update user
	if in.Username != nil {
		if err := user.UpdateUsername(*in.Username); err != nil {
			return nil, err
		}
	}

	// update user
	if err := uc.userRepository.Update(c, user); err != nil {
		return nil, err
	}

	// change User to OutputUser
	out := uc.userToUserOutput(user)

	return out, nil
}

func (uc *userUseCase) FindByID(c context.Context, in *InputFindByIDUser) (*OutputUser, error) {
	// find user
	user, err := uc.userRepository.FindByID(c, in.ID)
	if err != nil {
		return nil, err
	}

	// change User to OutputUser
	out := uc.userToUserOutput(user)

	return out, nil
}

func (uc *userUseCase) FindByUsername(c context.Context, in *InputFindByUsernameUser) (*OutputUser, error) {
	// find user
	user, err := uc.userRepository.FindByUsername(c, in.Username)
	if err != nil {
		return nil, err
	}

	// change User to OutputUser
	out := uc.userToUserOutput(user)

	return out, nil
}

// helper.
func (uc *userUseCase) userToUserOutput(user *entity.User) *OutputUser {
	return &OutputUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
