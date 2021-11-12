package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	str_page := c.Query("page")
	page, _ := strconv.Atoi(str_page)

	if page > 0 {
		result = (page - 1)
	}

	return result
}
