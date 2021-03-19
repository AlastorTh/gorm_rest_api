package routes

import (
	"github.com/AlastorTh/gorm_rest_api/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	u1 := r.Group("/admin")
	{
		u1.GET("/getads", controllers.GetAds)
		u1.POST("/createad", controllers.CreateAd)
		u1.GET("/get/:id", controllers.GetAdByID)
	}

	return r
}
