package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Kayky", "Kayky@.com", "123123")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Kayky", user.Name)
	assert.Equal(t, "Kayky@.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Kayky", "Kayky@.com", "123123")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassoword("123123"))
	assert.False(t, user.ValidatePassoword("123123123"))
	assert.NotEqual(t, "123123", user.Password)
}
