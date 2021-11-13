package services

import (
	m "github.com/ahmed-saleh/playbook/models"
)

type User struct {
	Id           string
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

//add params to listing later
func (u *User) List() ([]*m.User, error) {
	dat, err := m.GetUsers(u.PageNum, 3, u.getMaps())
	if err != nil {
		return nil, err
	}

	return dat, nil
}

//add map info
func (u *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	return maps
}

func (u *User) DeleteUser() error {
	if err := m.DeleteUser(u.Id); err != nil {
		return err
	}
	return nil
}
