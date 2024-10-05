package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required" gorm:"unique;not null"`
	Password string    `json:"password" binding:"required"`
}
