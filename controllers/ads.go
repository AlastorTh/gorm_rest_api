package controllers

import (
	"net/http"

	"github.com/AlastorTh/gorm_rest_api/configs"
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

func CreateAd(c *gin.Context) {
	var input models.CreateAdInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ad := models.Ad{Name: input.Name, Descript: input.Descript,
		Price: input.Price, MainPic: input.MainPic, OtherPics: input.OtherPics}
	ad.Validate()
	configs.DB.Create(&ad)
	c.JSON(http.StatusOK, gin.H{"ads": ad})
}

func GetAdByID(c *gin.Context) {
	var ad models.Ad
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&ad).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ads": ad})
}
