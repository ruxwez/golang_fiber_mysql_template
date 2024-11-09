package authInfrastructure

import (
	"fmt"

	"api.ruxwez.dev/adapters"
	"api.ruxwez.dev/adapters/mysql/models"
	jwtLibGen "api.ruxwez.dev/libs/jwt/generators"
)

func CreateAndRegisterAuthToken(user_id uint) string {

	jwt, err := jwtLibGen.GenerateAuthToken(fmt.Sprint(user_id))

	if err != nil {
		return ""
	}

	err = adapters.MySQL.Create(models.AuthToken{
		Token: jwt,
	}).Error

	if err != nil {
		return ""
	}

	return jwt
}