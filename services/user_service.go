package services

import (
	"fmt"

	m "github.com/ahmed-saleh/playbook/models"
)

func CreateUser(d []string) {
	fmt.Println(d)
}

type User struct {
	Display_name string
	Email        string
	PageNum      int
	PageSize     int
}

func (u *User) AddUser() error {
	user := map[string]interface{}{
		"email":        u.Email,
		"display_name": u.Display_name,
	}

	if err := m.AddUser(user); err != nil {
		return err
	}

	return nil
}
