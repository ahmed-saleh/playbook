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

func AddUser(data map[string]interface{}) error {
	user := &User{
		Email:        data["email"].(string),
		Display_name: data["display_name"].(string),
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
