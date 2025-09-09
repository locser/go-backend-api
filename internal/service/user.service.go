package serivce

import "go-backend/internal/repo"

type UserService struct {
	userRepo *repo.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo:  repo.NewUserRepository(),
	}
}

func (us *UserService) GetDetailUser() string {
	return us.userRepo.GetDetailUser()
}
