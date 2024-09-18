package validator

import (
	"gorm.io/gorm"
)

func Query(query *gorm.DB) (exists bool, err error) {
	if query.Error != nil {
		return false, query.Error
	}

	var result int64
	if query.Count(&result); result == 0 {
		return false, nil
	}

	return true, nil
}
