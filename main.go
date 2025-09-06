package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/davi-sant/househub-go/config"
	"github.com/davi-sant/househub-go/controllers"
	"github.com/davi-sant/househub-go/repositories"
	"github.com/davi-sant/househub-go/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("No .env file found")
	}

	ginMode := os.Getenv("GIN_MODE")

	config.DBConnection()
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(config.DB)

	config.RunMigrations("migrations/scheme.sql")

	gin.SetMode(ginMode)
	r := gin.Default()

	r.GET("/halph", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{})
	})
	port := ":3001"

	recordRepository := repositories.NewRecordRepository(config.DB)
	recordService := services.NewRecordService(recordRepository)
	recordController := controllers.NewRecordController(recordService)

	r.POST("/api/v1/registros", recordController.Create)
	r.GET("/api/v1/registros", recordController.FindAll)
	r.GET("/api/v1/:id/registros", recordController.FindById)
	r.PUT("/api/v1/:id/registros", recordController.Update)

	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
