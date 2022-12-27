//+ wireinject

package main

import (
	"di_sample/repo"
	"di_sample/service"

	"github.com/google/wire"
)

func InitializeUserService() *service.UserService {
	wire.Build(service.NewUserServiceSingleton, repo.NewUserRepo)
	return nil
}
