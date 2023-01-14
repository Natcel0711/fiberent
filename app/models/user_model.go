package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Username     string    `db:"username" json:"username" validate:"required"`
	Email        string    `db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
}
