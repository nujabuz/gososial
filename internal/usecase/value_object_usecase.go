package usecase

import "time"

// UserUseCase value objects

type InputCreateUser struct {
	Username string
	Password string
}

type InputUpdateUser struct {
	ID       int64
	Username *string
}

type InputFindByIDUser struct {
	ID int64
}

type InputFindByUsernameUser struct {
	Username string
}

type InputDeleteUser struct {
	ID int64
}

type OutputUser struct {
	ID        int64
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PostUseCase value objects

type InputCreatePost struct {
	UserID int64
	Body   string
}

type InputUpdatePost struct {
	UserID int64
	Body   *string
}

type InputDeletePost struct {
	ID     int64
	UserID int64
}

type InputFindByIDPost struct {
	ID int64
}

type InputFindByUserIDPost struct {
	UserID int64
}

type OutputPost struct {
	ID     int64
	UserID int64
	Body   string

	CreatedAt time.Time
	UpdatedAt time.Time
}
