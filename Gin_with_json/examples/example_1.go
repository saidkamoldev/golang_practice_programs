package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Foydalanuvchi uchun struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var newUser User

		// JSONni Go strcut ga dekodlash c.ShouldBindJSON struckga kelgan malumotlarni bog'laydi
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// qabul qilingan malumotni qaytarish
		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"user":    newUser,
		})
	})

	r.GET("/user", func(c *gin.Context) {
		var newUser User
		c.JSON(http.StatusOK, newUser)
	})

	r.Run(":8080")
}
