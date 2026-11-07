package entities

import (
	"fmt"
	"time"
)

type LikePost struct {
	ID        int64
	PostID    int64
	UserID    int64
	CreatedAt time.Time
}

func NewLikePost(postID int64, userID int64) *LikePost {
	likePost := &LikePost{
		UserID:    userID,
		PostID:    postID,
		CreatedAt: time.Now(),
	}

	return likePost
}

func (lk *LikePost) Validate() error {
	// validate post id
	if lk.PostID == 0 {
		return fmt.Errorf("%w: post id must not be empty", ErrValidation)
	}

	// validate user id
	if lk.UserID == 0 {
		return fmt.Errorf("%w: user id must not be empty", ErrValidation)
	}

	return nil
}
