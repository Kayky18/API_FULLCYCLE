package entity

import "github.com/google/uuid"

/*
func NewId() *string {
	id := uuid.New().String()
	return &id
}*/

type ID = uuid.UUID

func NewId() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	u, err := uuid.Parse(s)
	return ID(u), err
}
