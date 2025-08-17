package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPagination(c *gin.Context) (int, int, int, int) {
	const DEFAULT_PAGE = 1
	const INSTANCES_PER_PAGE = 10
	var page int
	if v, err := strconv.Atoi(c.Query("page")); err == nil && v > 0 {
		page = v
	} else {
		page = DEFAULT_PAGE
	}

	offset := (page - 1) * INSTANCES_PER_PAGE
	limit := offset + INSTANCES_PER_PAGE

	return offset, limit, page, INSTANCES_PER_PAGE
}
