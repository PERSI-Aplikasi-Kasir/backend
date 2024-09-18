package service

import (
	"backend/internal/module/user/entity"
	"backend/internal/module/user/interfaces"
	"fmt"
)

var _ interfaces.UserService = &userService{}

type userService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUser(req *entity.UserReqByUUID) (*entity.UserGet, error) {
	user, found, err := s.repo.GetUser(req.UUID)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("pengguna tidak ditemukan")
	}

	return &entity.UserGet{
		UUID: user.UUID,
		Name: user.Name,
	}, nil
}
