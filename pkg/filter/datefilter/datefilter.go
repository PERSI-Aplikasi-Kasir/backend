package datefilter

import (
	"backend/common/consts"
	"fmt"

	"gorm.io/gorm"
)

type DateFilter struct {
	OrderBy consts.OrderBy `json:"order_by" form:"order_by"`
	Sort    consts.Sort    `json:"sort" form:"sort"`
}

func Build(query *gorm.DB, filter *DateFilter) {
	if filter.OrderBy != consts.ASC && filter.OrderBy != consts.DESC {
		return
	}

	if filter.Sort != consts.CreatedAt && filter.Sort != consts.UpdatedAt {
		return
	}

	query.Order(fmt.Sprintf("%s %s", filter.Sort, filter.OrderBy))
}
