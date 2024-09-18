package interfaces

import "backend/internal/module/user/entity"

type UserRepository interface {
	GetUser(id string) (*entity.User, bool, error)
}

type UserService interface {
	GetUser(req *entity.UserReqByUUID) (*entity.UserGet, error)
}
