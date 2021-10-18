package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ahmed-saleh/playbook/pkg/app"
	"github.com/ahmed-saleh/playbook/services"
)

type CreateUserForm struct {
	Email        string
	Display_name string
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

	user_service := services.User{
		Display_name: form.Display_name,
		Email:        form.Email,
	}

	if err := user_service.AddUser(); err != nil {
		appG.Response(http.StatusInternalServerError, 500, map[string]interface{}{})
	}
}
