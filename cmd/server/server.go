package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ahmed-saleh/playbook/config"
	"github.com/ahmed-saleh/playbook/routers"
)

func Start(settings *config.Server) {
	r := gin.New()
	//load APIs
	routers.InitRouter(r)

	readTimeout := settings.ReadTimeout
	writeTimeout := settings.WriteTimeout
	endPoint := fmt.Sprintf(":%d", settings.HttpPort)
	maxHeaderBytes := 1 << 20
	gin.SetMode(settings.RunMode)

	server := &http.Server{
		Handler:        r,
		Addr:           endPoint,
		MaxHeaderBytes: maxHeaderBytes,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
	}
	server.ListenAndServe()
}
