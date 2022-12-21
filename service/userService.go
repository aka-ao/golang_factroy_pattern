package service

import (
	"di_sample/domain"
	"di_sample/repo"
)

type UserService struct {
	repo repo.IUserRepository
}

func NewUserService(repo repo.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) All() ([]domain.User, error) {
	return s.repo.All()
}
