package singlesearch

import (
	"gorm.io/gorm"
)

func Build(query *gorm.DB, searchField string, searchValue *string) {
	if searchValue == nil || *searchValue == "" {
		return
	}

	field := "LOWER(" + searchField + ")"
	query.Where(field+" LIKE LOWER(?)", "%"+*searchValue+"%")
}
