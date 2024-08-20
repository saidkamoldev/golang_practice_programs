package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

func main() {

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		// Foydalanuvchi malumotlari
		user := User{
			ID:   1,
			Name: "Alice",
		}

		c.JSON(http.StatusOK, user)
	})

	r.Run(":8080")
}
