package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/restlifeness/fire-proxy.git/src/api/routes"
)

func main() {
	r := gin.Default()

	r.GET("/ping", routes.PingRoute)
	r.POST("/auth", routes.AuthUserRoute)
	r.GET("/proxies", routes.GetAllAliveProxies)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
