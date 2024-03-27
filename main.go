package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"tugas-kedua/models"
	"tugas-kedua/routes"
)

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=database dbname=hacktivtugas2 port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.Order{}, &models.Item{})

	router := gin.Default()
	routes.SetupRoutes(router, db)

	if err := router.Run(":8080"); err != nil {
		panic("Failed to start server")
	}
}
