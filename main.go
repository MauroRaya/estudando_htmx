package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Name     string
	Quantity int
}

func newItem(name string, quantity int) Item {
	return Item{
		Name:     name,
		Quantity: quantity,
	}
}

type Error struct {
	IsNameInvalid     bool
	IsQuantityInvalid bool
}

func newError(isNameInvalid, isQuantityInvalid bool) Error {
	return Error{
		IsNameInvalid:     isNameInvalid,
		IsQuantityInvalid: isQuantityInvalid,
	}
}

func main() {
	r := gin.Default()

	items := []Item{
		newItem("banana", 50),
		newItem("cake", 10),
		newItem("orange", 20),
	}

	r.LoadHTMLGlob("templates/*")

	r.GET("/items", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"Error": Error{},
			"Items": items,
		})
	})

	r.POST("/items", func(c *gin.Context) {
		name := c.PostForm("Name")
		quantity := c.PostForm("Quantity")

		//validação nome duplicado
		for i := range items {
			if items[i].Name == name {
				c.HTML(422, "form", gin.H{
					"Error": newError(true, false),
				})
				return
			}
		}

		//validação quantidade
		quantityInt, err := strconv.Atoi(quantity)

		if err != nil {
			c.HTML(422, "form", gin.H{
				"Error": newError(false, true),
			})
			return
		}

		newItem := Item{name, quantityInt}
		items = append(items, newItem)

		c.HTML(200, "form", gin.H{
			"Error": Error{},
		})

		fmt.Println(items)

		c.HTML(200, "oob-item", newItem)
	})

	r.Run(":80")
}
