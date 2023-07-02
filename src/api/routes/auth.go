package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restlifeness/fire-proxy.git/src/api/schemas"
	"github.com/restlifeness/fire-proxy.git/src/api/security"
	"github.com/restlifeness/fire-proxy.git/src/api/services"
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

	jwtToken, err := security.GenerateJWTToken(requestBody)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	responseToken := schemas.NewBearerToken(jwtToken)

	ctx.JSON(200, responseToken)
}
