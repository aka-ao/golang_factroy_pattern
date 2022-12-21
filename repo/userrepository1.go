package repo

import "di_sample/domain"

type userRepository1 struct{}

func (ur *userRepository1) All() (users []domain.User, err error) {
	users = append(users, domain.User{
		Name:  "test1",
		Email: "test@example.com",
	})

	return users, err
}

func NewUserRepo1() IUserRepository {
	return &userRepository1{}
}
