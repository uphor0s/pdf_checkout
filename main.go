package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Invoice struct {
	Buyer_name string         `json:"buyer_name"`
	Contents   []CheckOutItem `json:"contents"`
}

type CheckOutItem struct {
	Name  string  `json:"item_name"`
	Price float64 `json:"item_price"`
	Count int     `json:"item_count"`
}

func main() {
	r := gin.Default()
	r.GET("/invoice", func(c *gin.Context) {
		var requestBody Invoice

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		fmt.Println(requestBody.Buyer_name)

		c.File(generate_invoice(requestBody))
	})
	// r.Run("26.231.160.213:5567")
	r.Run("0.0.0.0:5567")
}

//GOOS=windows GOARCH=amd64 go build  -o tg_bot_invoice_microservice.exe
