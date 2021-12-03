/*
	User model
*/
package models

import (
	"github.com/google/uuid"
	"github.com/gookit/validate"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model
	Display_name string
	Email        string
	Password     string
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.Uuid = uuid.New().String()
	return nil
}

func AddUser(data map[string]interface{}) error {
	v := validate.Map(data)
	//TODO: build a helper and apply on middleware as well
	v.StringRule("email", "required|minLen:3")
	v.StringRule("display_name", "required|minLen:3")
	v.StringRule("password", "required|minLen:8")

	if v.Validate() {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), 8)
		if err != nil {
			return err
		}
		user := &User{
			Email:        data["email"].(string),
			Display_name: data["display_name"].(string),
			Password:     string(hashedPassword),
		}
		if err := DB.Create(&user).Error; err != nil {
			return err
		}
	} else {
		return v.Errors
	}

	return nil
}

func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
	var users []*User
	err := DB.Offset(pageNum).Limit(pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

func DeleteUser(Id string) error {

	if err := DB.Delete(&User{}, Id).Error; err != nil {
		return err
	}
	return nil
}
