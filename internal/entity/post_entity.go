package entity

import (
	"fmt"
	"time"
)

type Post struct {
	ID     int64
	UserID int64
	Body   string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPost(userID int64, body string) (*Post, error) {
	song := &Post{
		UserID:    userID,
		Body:      body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := song.validate()
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (p *Post) UpdateBody(body string) error {
	p.Body = body
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Post) CheckUserID(userID int64) error {
	if p.UserID != userID {
		return fmt.Errorf("%w: post user id is not the same as '%d'", ErrValidation, userID)
	}

	return p.validate()
}

func (p *Post) validate() error {
	// validate userID.
	if p.UserID == 0 {
		return fmt.Errorf("%w: user id must not be empty", ErrValidation)
	}

	return nil
}
