package routers

import (
	v1 "github.com/ahmed-saleh/playbook/routers/api/v1"

	a "github.com/ahmed-saleh/playbook/routers/api/auth"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("/auth")
	auth.Use()
	{
		auth.POST("/login", a.Login)
	}

	apiV1 := r.Group("/api/v1")
	apiV1.Use()
	{
		apiV1.GET("/user", v1.ListUsers)
		apiV1.GET("/user/:id", v1.GetUser)
		apiV1.POST("/user", v1.CreateUser)
		apiV1.DELETE("/user/:id", v1.DeleteUser)
	}
}
