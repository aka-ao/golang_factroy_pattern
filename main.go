package main

import "fmt"

type User struct {
	Name  string
	Email string
}

type UserRepository interface {
	All() (users []User, err error)
}

type userRepository1 struct{}

func (ur *userRepository1) All() (users []User, err error) {
	users = append(users, User{
		Name:  "test1",
		Email: "test@example.com",
	})

	return users, err
}

func NewUserRepo1() UserRepository {
	return &userRepository1{}
}

type userRepository2 struct{}

func (ur *userRepository2) All() (users []User, err error) {
	users = append(users, User{
		Name:  "test",
		Email: "test@example.com",
	})

	users = append(users, User{
		Name:  "test2",
		Email: "test2@example.com",
	})

	return users, err
}

func NewUserRepo2() UserRepository {
	return &userRepository2{}
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) All() ([]User, error) {
	return s.repo.All()
}

func main() {
	userService1 := InitializeUserService1()
	users, _ := userService1.All()
	fmt.Println(users)

	userService2 := InitializeUserService2()
	users, _ = userService2.All()
	fmt.Println(users)
}
