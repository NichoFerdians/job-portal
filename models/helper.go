package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		// order := c.DefaultQuery("order_by", "id asc")

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize) //.Order(order)
	}
}
