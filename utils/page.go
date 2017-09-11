package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DefaultPageSizeStr = "100"
	DefaultPageSize    = 100

	DefaultFirstPageStr = "1"
	DefaultFirstPage    = 1
)

// first page: 1
func GetPaginationParams(c *gin.Context, maxPageSize int) (page, pageSize int) {
	pagestr := c.DefaultQuery("page", DefaultFirstPageStr)
	page, err := strconv.Atoi(pagestr)
	if err != nil || page <= 0 {
		page = DefaultFirstPage
	}

	pageSizestr := c.DefaultQuery("pageSize", DefaultPageSizeStr)
	pageSize, err = strconv.Atoi(pageSizestr)
	if err != nil || pageSize <= 0 || pageSize > maxPageSize {
		pageSize = DefaultPageSize
	}

	return page, pageSize
}
