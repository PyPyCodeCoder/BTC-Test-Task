# BTC-Test-Task

## Description
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

## Usage
To start the application via docker, use the following commands:

docker build -t test .

docker run -p 8080:8080 test

By default, the application will listen on port 8080.

## API Endpoints
The following API endpoints are available:

GET /api/rate: Retrieves the current exchange rate of Bitcoin to Ukrainian Hryvnia.

POST /api/subscribe: Subscribes an email address to receive exchange rate change notifications.

POST /api/sendEmails: Sends an email to all subscribed users with the current exchange rate.

### For any questions you can leave me a message:
  * @I3314 - telegram 
  * levvalos@gmail.com - email
