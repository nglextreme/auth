package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRef := gin.Default()

	ginRef.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ginRef.Run()
}
