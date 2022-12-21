// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func InitializeUserService1() *UserService {
	userRepository := NewUserRepo1()
	userService := NewUserService(userRepository)
	return userService
}

func InitializeUserService2() *UserService {
	userRepository := NewUserRepo2()
	userService := NewUserService(userRepository)
	return userService
}