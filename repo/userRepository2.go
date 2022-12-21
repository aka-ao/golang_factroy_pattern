package repo

import "di_sample/domain"

type userRepository2 struct{}

func (ur *userRepository2) All() (users []domain.User, err error) {
	users = append(users, domain.User{
		Name:  "test",
		Email: "test@example.com",
	})

	users = append(users, domain.User{
		Name:  "test2",
		Email: "test2@example.com",
	})

	return users, err
}

func NewUserRepo2() IUserRepository {
	return &userRepository2{}
}
