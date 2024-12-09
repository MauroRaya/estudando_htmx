package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Name     string
	Quantity int
}

func main() {
	r := gin.Default()

	items := []Item{
		{Name: "banana", Quantity: 50},
		{Name: "cake", Quantity: 12},
		{Name: "orange", Quantity: 20},
	}

	r.LoadHTMLGlob("templates/*")

	r.GET("/items", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"Items": items,
		})
	})

	r.POST("/items", func(c *gin.Context) {
		name := c.PostForm("Name")
		quantity := c.PostForm("Quantity")

		quantityInt, _ := strconv.Atoi(quantity)

		newItem := Item{name, quantityInt}

		items = append(items, newItem)

		c.HTML(200, "items", gin.H{
			"Items": items,
		})
	})

	r.Run(":80")
}
