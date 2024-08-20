package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	r := gin.Default()

	// Bir nechta foydalanuvchilarni saqlash uchun kesim
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	r.GET("/users/:id", func(c *gin.Context) {

		id := c.Param("id")
		// Foydalanuvchini ro'yhatdan qidirish
		for _, user := range users {
			if strconv.Itoa(user.ID) == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
	})

	r.Run(":8080")
}
