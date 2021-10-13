package main

import (
	"fmt"
	"net/http"

	"github.com/ahmed-saleh/playbook/routers"
)

func init() {

}

func main() {
	routersInit := routers.InitRouter()
	maxHeaderBytes := 1 << 20
	endPoint := fmt.Sprintf(":%d", 8080) //ini this

	server := &http.Server{
		Handler:        routersInit,
		Addr:           endPoint,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()
}
