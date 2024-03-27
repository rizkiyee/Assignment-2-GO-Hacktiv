package routes

import (
	"tugas-kedua/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/orders", func(c *gin.Context) {
		controllers.GetAllOrders(c, db)
	})
	router.POST("/orders", func(c *gin.Context) {
		controllers.CreateOrder(c, db)
	})
	router.PUT("/orders/:orderId", func(c *gin.Context) {
		controllers.UpdateOrder(c, db)
	})
	router.DELETE("/orders/:orderId", func(c *gin.Context) {
		controllers.DeleteOrder(c, db)
	})
}
