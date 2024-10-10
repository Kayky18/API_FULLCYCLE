package database

import (
	"github.com/Kayky18/API_FULLCYCLE/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
