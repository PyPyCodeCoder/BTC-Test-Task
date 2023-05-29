package main

import (
	"TestTaskT/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.GET("/api/rate", handlers.GetExchangeRateHandler)
	router.POST("/api/subscribe", handlers.SubscribeEmailHandler)
	router.POST("/api/sendEmails", handlers.SendEmailsHandler)

	err = router.Run(":8080")
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
	}
}
