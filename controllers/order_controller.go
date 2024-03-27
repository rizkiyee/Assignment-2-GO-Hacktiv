package controllers

import (
	"net/http"

	"tugas-kedua/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllOrders(c *gin.Context, db *gorm.DB) {
	var orders []models.Order
	db.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context, db *gorm.DB) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&order)
	c.JSON(http.StatusCreated, order)
}

func UpdateOrder(c *gin.Context, db *gorm.DB) {
	orderId := c.Param("orderId")
	var order models.Order
	if err := db.Preload("Items").First(&order, orderId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var updatedOrder models.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.CustomerName = updatedOrder.CustomerName
	order.OrderedAt = updatedOrder.OrderedAt

	for _, item := range order.Items {
		db.Delete(&item)
	}
	order.Items = updatedOrder.Items

	db.Save(&order)
	response := gin.H{
		"customer_name": order.CustomerName,
		"ordered_at":    order.OrderedAt,
		"items":         order.Items,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteOrder(c *gin.Context, db *gorm.DB) {
	orderId := c.Param("orderId")
	var order models.Order
	if err := db.Preload("Items").First(&order, orderId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	for _, item := range order.Items {
		db.Delete(&item)
	}

	db.Delete(&order)

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
