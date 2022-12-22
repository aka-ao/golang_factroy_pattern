//+ wireinject

package main

import (
	"di_sample/repo"
	"di_sample/service"

	"github.com/google/wire"
)

func InitializeUserService1() *service.UserService {
	wire.Build(service.NewUserService, repo.NewUserRepo1)
	return nil
}

func InitializeUserService2() *service.UserService {
	wire.Build(service.NewUserService, repo.NewUserRepo2)
	return nil
}

func InitializeUserServiceSingleton() *service.UserService {
	wire.Build(service.NewUserServiceSingleton, repo.NewUserRepo1)
	return nil
}
