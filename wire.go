//+ wireinject

package main

import (
	"di_sample/repo"
	"di_sample/service"
)

func initializeUserService() *service.UserService {
	iUserRepository := repo.NewUserRepo()
	userService := service.NewUserServiceSingleton(iUserRepository)
	return userService
}
