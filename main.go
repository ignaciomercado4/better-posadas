package main

import (
	"better-posadas/database"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.Use(cors.Default())

	db := database.ConnectDatabase()
	database.MigrateModels(db)

	reportHandler := ReportHandler{DB: db}

	r.LoadHTMLGlob("templates/*")

	r.GET("/reports", reportHandler.GetReports)

	r.POST("/reports/create", reportHandler.CreateReport)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	r.Run(":" + string(PORT))
}
