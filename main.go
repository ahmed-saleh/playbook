package main

import (
	"fmt"
	"net/http"

	"github.com/ahmed-saleh/playbook/config"
	"github.com/ahmed-saleh/playbook/models"
	"github.com/ahmed-saleh/playbook/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Setup()
	models.Setup()
}

func main() {
	gin.SetMode(config.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Handler:        routersInit,
		Addr:           endPoint,
		MaxHeaderBytes: maxHeaderBytes,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
	}
	server.ListenAndServe()
}
