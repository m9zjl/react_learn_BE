package service

import (
	"server/pkg/domain/entity"
	"server/pkg/repo"
)

type UserService struct {
	userRepo repo.IUserRepo
}

func NewUserService(userRepo repo.IUserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	return s.userRepo.ByEmail(email)
}
