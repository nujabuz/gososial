// Package entity layer
package entity

import (
	"fmt"
	"time"
)

var (
	MaxCharUsername = 16
	MinCharUsername = 4
	MaxCharPassword = 24
	MinCharPassword = 8
)

type User struct {
	ID       int64
	Username string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username string, password string) (*User, error) {
	user := &User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := user.validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) UpdateUsername(username string) error {
	u.Username = username
	u.UpdatedAt = time.Now()

	return u.validate()
}

func (u *User) UpdatePassword(password string) error {
	u.Password = password
	u.UpdatedAt = time.Now()

	return u.validate()
}

func (u *User) validate() error {
	lenUsername := len(u.Username)
	lenPassword := len(u.Password)

	// validate username.
	if u.Username == "" {
		return fmt.Errorf("%w: username must not be empty", ErrValidation)
	}

	if lenUsername > MaxCharUsername {
		return fmt.Errorf("%w: username must not more than %d", ErrValidation, MaxCharUsername)
	}

	if lenUsername < MinCharUsername {
		return fmt.Errorf("%w: username must not less than %d", ErrValidation, MinCharUsername)
	}

	// validate password.
	if u.Password == "" {
		return fmt.Errorf("%w: password must not be empty", ErrValidation)
	}

	if lenPassword > MaxCharPassword {
		return fmt.Errorf("%w: password must not more than %d", ErrValidation, MaxCharPassword)
	}

	if lenPassword < MinCharPassword {
		return fmt.Errorf("%w: password must not less than %d", ErrValidation, MinCharPassword)
	}

	return nil
}
