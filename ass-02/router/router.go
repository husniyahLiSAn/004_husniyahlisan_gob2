package router

import (
	"ass-02/config"
	"ass-02/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	db := config.GetDB()
	route := &controllers.Database{Connect: db}

	orderRouter := r.Group("/orders")
	{
		orderRouter.POST("/", route.CreateOrder)
		orderRouter.GET("/", route.GetAllOrders)
		orderRouter.PUT("/:orderId", route.UpdateOrder)
		orderRouter.DELETE("/:orderId", route.DeleteOrder)
	}

	return r
}
