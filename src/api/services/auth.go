package services

import (
	"fmt"

	"github.com/restlifeness/fire-proxy.git/src/api/repo"
	"github.com/restlifeness/fire-proxy.git/src/api/schemas"
	"github.com/restlifeness/fire-proxy.git/src/api/security"
)

func AuthUser(user schemas.RequestAuthForm) (bool, error) {
	dbUser, err := repo.GetUserByUsername(user.Username)
	if err != nil {
		return false, err
	}
	fmt.Println("hashed ass", dbUser.HashedPassword)
	fmt.Println("password", user.Password)
	result, err := security.ComparePasswords(user.Password, dbUser.HashedPassword)
	return result, err
}
