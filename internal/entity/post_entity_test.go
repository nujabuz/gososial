package entity_test

import (
	"testing"

	"github.com/nujabuz/gososial/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestPost_NewPost_Success(t *testing.T) {
	post, err := entity.NewPost(1, "Hello World")

	assert.Nil(t, err)
	assert.NotNil(t, post)

	assert.Equal(t, int64(1), post.UserID)
	assert.Equal(t, "Hello World", post.Body)
}

func TestPost_NewPost_Failed_EmptyUserID(t *testing.T) {
	post, err := entity.NewPost(0, "Hello World")

	assert.Nil(t, post)
	assert.NotNil(t, err)
}

func TestPost_NewPost_CheckUserID(t *testing.T) {
	t.Run("success check user id", func(t *testing.T) {
		post, err := entity.NewPost(1, "Hello World")

		assert.Nil(t, err)
		assert.NotNil(t, post)

		assert.Equal(t, int64(1), post.UserID)
		assert.Equal(t, "Hello World", post.Body)

		err = post.CheckUserID(1)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), post.UserID)
	})

	t.Run("failed check user id", func(t *testing.T) {
		post, err := entity.NewPost(1, "Hello World")

		assert.Nil(t, err)
		assert.NotNil(t, post)

		assert.Equal(t, int64(1), post.UserID)
		assert.Equal(t, "Hello World", post.Body)

		err = post.CheckUserID(2)
		assert.NotNil(t, err)
	})
}
