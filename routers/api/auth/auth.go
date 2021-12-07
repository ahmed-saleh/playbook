package auth

import (
	"net/http"

	"github.com/ahmed-saleh/playbook/models"
	"github.com/ahmed-saleh/playbook/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	data := make(map[string]interface{})
	data["email"] = c.PostForm("email")
	data["password"] = c.PostForm("password")

	v := validate.Map(data)
	v.StringRule("email", "required")
	v.StringRule("password", "required")

	if v.Validate() {
		user, err := models.FindByEmail(data["password"].(string))
		if err != nil {
			appG.Response(http.StatusBadRequest, 422, "Authentication failed")
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"].(string)))
		if err != nil {
			appG.Response(http.StatusBadRequest, 422, "Authentication failed")
		}
		appG.Response(http.StatusBadRequest, 422, "Authentication failed")
	} else {
		appG.Response(http.StatusBadRequest, 400, map[string]string{
			"error": v.Errors.Error(),
		})
	}

	appG.Response(http.StatusOK, 200, map[string]string{
		"token": "token",
	})
}
