package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ahmed-saleh/playbook/pkg/app"
	"github.com/ahmed-saleh/playbook/pkg/utils"
	"github.com/ahmed-saleh/playbook/services"
)

type CreateUserForm struct {
	Email        string
	Display_name string
	Password     string
}

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, 22, map[string]interface{}{
		"lists": []string{"first", "second"},
		"total": 2,
	})
}

func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var form CreateUserForm
	_ = c.Bind(&form)

	userService := services.User{
		Display_name: form.Display_name,
		Email:        form.Email,
		Password:     form.Password,
	}

	if err := userService.AddUser(); err != nil {
		appG.Response(http.StatusInternalServerError, 500, err)
		appG.C.Abort()
		return
	}
	appG.Response(http.StatusCreated, 201, "User created successfully")
}

func ListUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	userService := services.User{
		PageNum: utils.GetPage(c),
	}

	data, _ := userService.List()
	appG.Response(http.StatusOK, 200, data)

}

func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	params, _ := c.Params.Get("id")

	userService := services.User{
		Id: params,
	}

	if err := userService.DeleteUser(); err != nil {
		appG.Response(http.StatusInternalServerError, 422, err)
	}

	appG.Response(http.StatusOK, 200, "User was deleted successfully.")

}
