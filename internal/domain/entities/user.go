// Package entities
package entities

import (
	"fmt"
	"time"
)

var (
	MaxCharUsername = 16
	MaxCharPassword = 24
	MinCharUsername = 3
	MinCharPassword = 8
)

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username string, password string) *User {
	user := &User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user
}

func (u *User) UpdateUsername(username string) error {
	u.Username = username
	u.UpdatedAt = time.Now()

	return u.Validate()
}

func (u *User) Validate() error {
	// validatee username
	if u.Username == "" {
		return fmt.Errorf("%w: username must not be empty", ErrValidation)
	}

	if len(u.Username) > MaxCharUsername {
		return fmt.Errorf("%w: length username must not more than %d", ErrValidation, MaxCharUsername)
	}

	if len(u.Username) < MinCharUsername {
		return fmt.Errorf("%w: length username must not less than %d", ErrValidation, MinCharUsername)
	}

	// validate username
	if u.Password == "" {
		return fmt.Errorf("%w: password must not be empty", ErrValidation)
	}

	if len(u.Password) > MaxCharPassword {
		return fmt.Errorf("%w: length password must not more than %d", ErrValidation, MaxCharPassword)
	}

	if len(u.Password) < MinCharPassword {
		return fmt.Errorf("%w: length password must not less than %d", ErrValidation, MinCharPassword)
	}

	return nil
}
