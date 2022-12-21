package repo

import "di_sample/domain"

type IUserRepository interface {
	All() (users []domain.User, err error)
}
