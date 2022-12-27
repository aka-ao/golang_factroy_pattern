package main

import (
	"log"
)

func main() {
	userService := InitializeUserService()
	users, _ := userService.All()
	for _, user := range users {
		log.Println(user.Pay(1000))
	}
}
