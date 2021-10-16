/*
	User model
*/
package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Model
	Display_name string
	Email        string
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.Uuid = uuid.New().String()
	return nil
}
