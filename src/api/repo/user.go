package repo

import (
	"github.com/restlifeness/fire-proxy.git/src/database"
)

func GetUserByUsername(username string) (db.User, error) {
	var user database.User
	err := db.First("username = ?", username).First(&user).Error
	return user, err
}
