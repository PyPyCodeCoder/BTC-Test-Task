package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func GetExchangeRateHandler(c *gin.Context) {
	exchangeRate, err := GetBitcoinExchangeRate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rate": exchangeRate})
}

func GetBitcoinExchangeRate() (float64, error) {
	url := os.Getenv("url")

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	rates := data["data"].(map[string]interface{})["rates"].(map[string]interface{})
	rate := rates["UAH"].(string)

	var exchangeRate float64
	_, err = fmt.Sscan(rate, &exchangeRate)
	if err != nil {
		return 0, err
	}

	return exchangeRate, nil
}
