package stringfilter

import (
	"backend/common/consts"
	"fmt"

	"gorm.io/gorm"
)

type StringFilter struct {
	OrderBy consts.OrderBy `json:"order_by" form:"order_by"`
	Sort    string         `json:"sort" form:"sort"`
}

func Build(query *gorm.DB, fieldName string, filter *StringFilter) {
	if filter.OrderBy != consts.ASC && filter.OrderBy != consts.DESC {
		return
	}

	if filter.Sort != fieldName && filter.OrderBy != consts.ASC && filter.OrderBy != consts.DESC {
		return
	}

	query.Order(fmt.Sprintf("%s %s", fieldName, filter.OrderBy))
}
