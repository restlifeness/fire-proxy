package repo

import (
	"github.com/restlifeness/fire-proxy.git/src/database"
)

func GetAliveProxiesConnections() ([]database.Proxy, error) {
	var proxy []database.Proxy

	db := database.ConnectToDatabase()
	err := db.Where("still_alive = ?", true).Find(&proxy).Error
	return proxy, err
}
