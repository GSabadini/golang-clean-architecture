package entity

import (
	gouuid "github.com/google/uuid"
)

func NewUUID() string {
	return gouuid.New().String()
}
