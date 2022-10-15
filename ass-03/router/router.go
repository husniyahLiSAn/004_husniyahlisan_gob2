package router

import (
	"ass-03/config"
	"ass-03/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	db := config.GetDB()
	route := &controllers.Database{Connect: db}

	orderRouter := r.Group("/weather")
	{
		orderRouter.POST("/", route.CreateData)
		orderRouter.POST("/update", route.Update)
		orderRouter.GET("/", route.GetLatestData)
	}

	return r
}
