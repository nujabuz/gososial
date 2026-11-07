package entities

import (
	"fmt"
	"time"
)

type Post struct {
	ID        int64
	UserID    int64
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPost(userID int64, body string) *Post {
	post := &Post{
		UserID:    userID,
		Body:      body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return post
}

func (p *Post) Validate() error {
	// validate userID
	if p.UserID == 0 {
		return fmt.Errorf("%w: user id must not be empty", ErrValidation)
	}

	return nil
}
