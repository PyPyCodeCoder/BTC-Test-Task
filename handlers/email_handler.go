package handlers

import (
	"TestTaskT/utils"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SubscribeEmailHandler(c *gin.Context) {
	email := c.PostForm("email")

	if !utils.ValidateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if utils.EmailExists(email) {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	err := utils.SaveEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email subscribed"})
}

func SendEmailsHandler(c *gin.Context) {
	exchangeRate, err := GetBitcoinExchangeRate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filePath := os.Getenv("filePath")
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		email := scanner.Text()
		err := utils.SendEmail(email, exchangeRate)
		if err != nil {
			fmt.Printf("Failed to send email to %s: %s\n", email, err.Error())
		} else {
			fmt.Printf("Email sent successfully to %s\n", email)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Emails sent"})
}
