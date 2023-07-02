package routes

import (
	"github.com/restlifeness/fire-proxy.git/src/api/schemas"
	"github.com/restlifeness/fire-proxy.git/src/api/services"

	"github.com/gin-gonic/gin"
)

func AuthUserRoute(ctx *gin.Context) {
	var requestBody schemas.RequestAuthForm

	if err := ctx.Bind(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := services.AuthUser(requestBody)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if !result {
		ctx.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	simpleToken := "i_love_fire"
	responseToken := schemas.NewBearerToken(simpleToken)

	ctx.JSON(200, responseToken)
}
