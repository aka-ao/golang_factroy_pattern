//+ wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeUserService1() *UserService {
	wire.Build(NewUserService, NewUserRepo1)
	return nil
}

func InitializeUserService2() *UserService {
	wire.Build(NewUserService, NewUserRepo2)
	return nil
}
