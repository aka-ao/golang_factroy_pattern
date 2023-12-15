package main

import (
	"log"
)

func main() {
	userService := initializeUserService()
	users, _ := userService.All()
	for _, user := range users {
		log.Println(user.Pay(1000))
	}
}
