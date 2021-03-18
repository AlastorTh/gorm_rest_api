package main

import (
	"github.com/AlastorTh/gorm_rest_api/configs"
	"github.com/AlastorTh/gorm_rest_api/models"
	"github.com/AlastorTh/gorm_rest_api/routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	configs.DB, err = gorm.Open("postgres",
		"host=db port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer configs.DB.Close()
	configs.DB.AutoMigrate(&models.Ad{})

	r := routes.SetUpRouter()
	r.Run(":8080")
}
