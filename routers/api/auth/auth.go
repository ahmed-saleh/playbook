package auth

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"golang.org/x/crypto/bcrypt"

	"github.com/ahmed-saleh/playbook/models"
)

type LoginForm struct {
	Email    string
	Password string
}

func Login(c *gin.Context) (interface{}, error) {
	var form LoginForm
	_ = c.Bind(&form)

	data := make(map[string]interface{})
	data["email"] = form.Email
	data["password"] = form.Password
	v := validate.Map(data)
	v.StringRule("email", "required")
	v.StringRule("password", "required")

	if v.Validate() {
		user, err := models.FindByEmail(data["email"].(string))
		fmt.Println(user)

		if err != nil {
			return "", err
		}
		if user == (models.User{}) {
			//TODO: log no user login
			return "", errors.New("authentication faild")
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"].(string)))
		if err != nil {
			return "", errors.New("authentication faild")
		}
		return "token", nil
	} else {
		return "", v.Errors
	}
}
