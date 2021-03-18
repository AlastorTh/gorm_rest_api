package controllers

import (
	"net/http"

	"github.com/AlastorTh/gorm_rest_api/models"
	"github.com/AlastorTh/gorm_rest_api/repo"
	"github.com/AlastorTh/gorm_rest_api/utils"
	"github.com/gin-gonic/gin"
)

func GetAds(c *gin.Context) {
	pagination := utils.GeneratePagination(c)
	var ad models.Ad
	adLists, err := repo.GetAllAds(&ad, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": adLists,
	})

}
