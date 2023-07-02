package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restlifeness/fire-proxy.git/src/api/repo"
)

func GetAllAliveProxies(ctx *gin.Context) {
	proxies, err := repo.GetAliveProxiesConnections()

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, proxies)
}
