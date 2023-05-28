# BTC-Test-Task

This repository contains a Go application that implements an API service for managing exchange rates and email subscriptions. The service provides the following functionalities:

Get the current exchange rate of Bitcoin (BTC) to Ukrainian Hryvnia (UAH).
Subscribe an email address to receive information about exchange rate changes.
Send an email to all subscribed users with the current exchange rate.

## Prerequisites
Before running the application, make sure you have the following installed:

Go programming language (version 1.16 or higher)
github.com/gin-gonic/gin package
github.com/joho/godotenv package

## Installation
Clone this repository to your local machine.
Change to the project directory.
Install the required dependencies.

## Configuration
The application uses environment variables for configuration. Create a .env file in the project directory and provide the following variables:
API_PORT=8080

## Usage
To start the application, run the following command:
go run main.go

By default, the application will listen on port 8080. You can change the port by modifying the API_PORT variable in the .env file.

## API Endpoints
The following API endpoints are available:

GET /api/rate: Retrieves the current exchange rate of Bitcoin to Ukrainian Hryvnia.

POST /api/subscribe: Subscribes an email address to receive exchange rate change notifications.

POST /api/sendEmails: Sends an email to all subscribed users with the current exchange rate.
