package routes

import "github.com/gin-gonic/gin"

func PingRoute(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	message := "pong"

	messages := query["message"]
	if messages != nil {
		message = messages[0]
	}

	ctx.JSON(200, gin.H{
		"message": message,
	})
}
