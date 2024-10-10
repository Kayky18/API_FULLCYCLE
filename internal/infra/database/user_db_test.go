package database

import (
	"testing"

	"github.com/Kayky18/API_FULLCYCLE/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("test", "test@", "123123")

	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id  = ?", user.Id).Error
	assert.Nil(t, err)
	assert.Equal(t, user.Id, userFound.Id)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
