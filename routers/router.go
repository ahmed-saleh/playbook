package routers

import (
	v1 "github.com/ahmed-saleh/playbook/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	apiV1.Use()
	{
		apiV1.GET("/user", v1.GetUser)
		apiV1.POST("/user", v1.CreateUser)
	}

	return r
}
