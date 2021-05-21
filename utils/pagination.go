package utils

import (
	"strconv"

	"github.com/AlastorTh/gorm_rest_api/models"
	"github.com/gin-gonic/gin"
)

func GeneratePagination(c *gin.Context) models.Pagination {
	limit := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue

		}
	}

	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
