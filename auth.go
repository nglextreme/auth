package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nglxtreme/auth/aa"
	"nglxtreme/auth/config"
)

func main() {
	config.LoadProperties()
	config.ConnectToDatabase()

	ginRef := gin.Default()

	ginRef.GET("/ping", func(ctx *gin.Context) {
		aa.Authenticate()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ginRef.Run()
}
