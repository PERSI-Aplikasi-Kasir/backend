package entity

import (
	models "backend/common/model"

	"github.com/google/uuid"
)

type User struct {
	UUID uuid.UUID `gorm:"primaryKey;not null" validate:"required"`
	Name string    `gorm:"not null"`

	models.TimestampsSoftDelete
}

type UserReqByUUID struct {
	UUID string `uri:"uuid" binding:"required"`
}

type UserGet struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
