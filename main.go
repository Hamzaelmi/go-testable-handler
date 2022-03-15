package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	controller := PublicController{Database: "pong"}
	r.GET("/ping", controller.pong)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type PublicController struct {
	Database string
}

func (c *PublicController) pong(context *gin.Context) {

	context.JSON(200, gin.H{"msg": c.Database})
}
