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

func (s *UserService) AddUser(user *entity.User) (bool, error) {
	success, err := s.userRepo.Add(user)
	if err != nil {
		return success, err
	}
	//log.Info("add user email: %s :%s", user.Email, success)
	return success, err
}
