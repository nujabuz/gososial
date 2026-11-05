package entity_test

import (
	"testing"

	"github.com/nujabuz/gososial/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestUser_NewUser_Success(t *testing.T) {
	user, err := entity.NewUser("nujabuz", "nujabuz123")

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, "nujabuz", user.Username)
	assert.Equal(t, "nujabuz123", user.Password)
}

func TestUser_NewUser_Failed_EmptyUsername(t *testing.T) {
	// test username empty
	user, err := entity.NewUser("", "nujabuz123")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestUser_NewUser_Failed_UsernameMoreThanMaximum(t *testing.T) {
	// test username more than maximum
	user, err := entity.NewUser("nujabuz.nujabuz.n", "nujabuz123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUser_NewUser_Failed_UsernameLessThanMinimum(t *testing.T) {
	// test username less than minimum
	user, err := entity.NewUser("nuj", "nujabuz123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUser_NewUser_Failed_EmptyPassword(t *testing.T) {
	// test password empty
	user, err := entity.NewUser("nujabuz", "")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestUser_NewUser_Failed_PasswordMoreThanMaximum(t *testing.T) {
	// test password  more than maximum
	user, err := entity.NewUser("nujabuz", "1234567891234567891234567")

	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUser_NewUser_Failed_PasswordLessThanMinimum(t *testing.T) {
	// test username less than minimum
	user, err := entity.NewUser("nujabuz", "1234567")

	assert.Nil(t, user)
	assert.NotNil(t, err)
}
