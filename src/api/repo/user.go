package repo

import (
	"github.com/restlifeness/fire-proxy.git/src/database"
)

func GetUserByUsername(username string) (database.User, error) {
	var user database.User

	db := database.ConnectToDatabase()
	err := db.Model(&database.User{}).Where("username = ?", username).First(&user).Error
	return user, err
}
