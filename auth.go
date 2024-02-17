package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"nglxtreme/auth/aa"
	"nglxtreme/auth/config"
)

func main() {
	config.LoadProperties()
	config.ConnectToDatabase()

	ginRef := gin.Default()

	ginRef.GET("/authenticate", aa.HandleAuthenticateRequest)

	ginRef.Run(viper.GetString("port"))
}
