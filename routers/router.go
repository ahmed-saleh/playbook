package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/ahmed-saleh/playbook/routers/api"
	v1 "github.com/ahmed-saleh/playbook/routers/api/v1"
)

func InitRouter(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	m := api.Setup()

	auth := r.Group("/api/auth")
	auth.Use()
	{
		// auth.POST("/login", m.LoginHandler)
		auth.POST("/login", m.LoginHandler)

	}

	apiV1 := r.Group("/api/v1")
	apiV1.Use(m.MiddlewareFunc())
	{
		apiV1.GET("/user", v1.ListUsers)
		apiV1.GET("/user/:id", v1.GetUser)
		apiV1.POST("/user", v1.CreateUser)
		apiV1.DELETE("/user/:id", v1.DeleteUser)
	}
}
