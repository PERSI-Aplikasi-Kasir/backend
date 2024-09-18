package migrations

import (
	userEntity "backend/internal/module/user/entity"

	"gorm.io/gorm"
)

type MigratableModel interface {
	Migrate(db *gorm.DB) error
}

var modelList = []interface{}{
	&userEntity.User{},
}

func Migrate(db *gorm.DB) error {
	for _, model := range modelList {
		if err := db.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
	return nil
}
