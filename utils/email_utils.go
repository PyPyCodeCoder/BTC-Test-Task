package utils

import (
	"bufio"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
	"strings"
)

func ValidateEmail(email string) bool {
	return strings.Contains(email, "@")
}

func EmailExists(email string) bool {
	filePath := os.Getenv("filePath")
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == email {
			return true
		}
	}
	return false
}

func SaveEmail(email string) error {
	filePath := os.Getenv("filePath")
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := fmt.Fprintln(file, email); err != nil {
		return err
	}

	return nil
}

func SendEmail(email string, exchangeRate float64) error {
	smtpUsername := os.Getenv("smtpUsername")
	emailSubject := os.Getenv("emailSubject")
	emailBodyFormat := os.Getenv("emailBodyFormat")
	smtpServer := os.Getenv("smtpServer")
	smtpPortSTR := os.Getenv("smtpPort")
	smtpPort, err := strconv.Atoi(smtpPortSTR)
	smtpPassword := os.Getenv("smtpPassword")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", smtpUsername)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", emailSubject)
	mailer.SetBody("text/plain", fmt.Sprintf(emailBodyFormat, exchangeRate))

	dialer := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
