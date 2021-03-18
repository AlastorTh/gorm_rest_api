package repo

import (
	"github.com/AlastorTh/gorm_rest_api/configs"
	"github.com/AlastorTh/gorm_rest_api/models"
)

func GetAllAds(ad *models.Ad, pagination *models.Pagination) (*[]models.Ad, error) {
	var ads []models.Ad
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := configs.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Ad{}).Where(ad).Find(&ads)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &ads, nil
}
