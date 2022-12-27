package service

import (
	"di_sample/domain"
	"di_sample/repo"
	"log"
)

type UserService struct {
	repo repo.IUserRepository
}

var instance *UserService

func NewUserServiceSingleton(repo repo.IUserRepository) *UserService {
	if instance == nil {
		log.Println("create new service instance")
		instance = &UserService{
			repo: repo,
		}
	}
	return instance
}

func (s *UserService) All() ([]domain.User, error) {
	return s.repo.All()
}
