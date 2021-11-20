package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ahmed-saleh/playbook/config"
	"github.com/ahmed-saleh/playbook/routers"
)

func Start() {
	r := gin.New()
	//load APIs
	routers.InitRouter(r)

	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	gin.SetMode(config.ServerSetting.RunMode)

	server := &http.Server{
		Handler:        r,
		Addr:           endPoint,
		MaxHeaderBytes: maxHeaderBytes,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
	}
	server.ListenAndServe()
}
