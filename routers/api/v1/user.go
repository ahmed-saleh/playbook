package v1

import (
	"net/http"

	"github.com/ahmed-saleh/playbook/pkg/app"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, 22, map[string]interface{}{
		"lists": []string{"first", "second"},
		"total": 2,
	})
}
