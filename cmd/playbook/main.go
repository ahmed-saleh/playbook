package main

import (
	"github.com/gin-gonic/gin"
	"log"

    "github.com/elastic/go-elasticsearch/v8"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": callme(),
		})
	})
	r.Run()
}

func callme() string {
	es, _ := elasticsearch.NewDefaultClient()
    log.Println(elasticsearch.Version)
    log.Println(es.Info())
	return "test api!"
}
