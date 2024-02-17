package aa

import (
	"github.com/gin-gonic/gin"
)

func HandleAuthenticateRequest(context *gin.Context) {
	username := context.Request.Header.Get("username")
	password := context.Request.Header.Get("password")

	patron, error := Authenticate(username, password)

	if error != nil {
		context.JSON(error.Status, *error)
	} else {
		context.JSON(200, *patron)
	}
}
