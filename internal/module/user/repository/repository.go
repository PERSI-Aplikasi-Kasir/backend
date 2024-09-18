package repository

import (
	"backend/internal/module/user/entity"
	"backend/internal/module/user/interfaces"
	"backend/pkg/validator"

	"gorm.io/gorm"
)

var _ interfaces.UserRepository = &userRepository{}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUser(id string) (*entity.User, bool, error) {
	var user entity.User
	query := r.db.Model(&user).Where("uuid = ?", id)
	exists, err := validator.Query(query)
	if err != nil {
		return nil, false, err
	}

	if !exists {
		return nil, false, nil
	}

	query.First(&user)
	return &user, true, nil
}
