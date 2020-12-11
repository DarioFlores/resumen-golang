package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(logger)

	r.GET("/", funcionQueHaceGet)

	r.Run(":8080")
}

func logger(c *gin.Context) {
	log.Println("Request!")
}

func funcionQueHaceGet(c *gin.Context) {
	c.String(200, "ok")
}
