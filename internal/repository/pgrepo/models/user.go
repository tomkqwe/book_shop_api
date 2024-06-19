package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	Role     string    `json:"role" db:"role"`
}

func (u User) GetId() uuid.UUID {
	return u.ID
}
