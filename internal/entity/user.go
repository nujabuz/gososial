package entity

import (
	"time"
)

var (
	MaxCharUsername = 16
	MinCharUsername = 4
	MaxCharPassword = 24
	MinCharPassword = 8
)

type User struct {
	ID        int64
	Username  string
	Password  string
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

func (u *User) validate() error {
	lenUsername := len(u.Username)
	lenPassword := len(u.Password)

	if u.Username == "" {
		return ErrValidation
	}

	if lenUsername > MaxCharUsername {
		return ErrValidation
	}

	if lenUsername < MinCharUsername {
		return ErrValidation
	}

	if u.Password == "" {
		return ErrValidation
	}

	if lenPassword > MaxCharPassword {
		return ErrValidation
	}

	if lenPassword < MinCharPassword {
		return ErrValidation
	}

	return nil
}
