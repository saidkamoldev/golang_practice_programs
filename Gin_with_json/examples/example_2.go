package main

import (
	"net/http"
	// "os/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Foydalanuvchi uchun struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var newUser User

		// Json ni go strucktiga dekodlash
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//foydalanuvchiga ID berish va ro'yhatga qo'shish
		newUser.ID = len(users) + 1
		users = append(users, newUser)

		c.JSON(http.StatusOK, gin.H{
			"message": "User created succeessfully",
			"users":   newUser,
		})
	})

	r.GET("/user/:id", func(c *gin.Context) {

		// url parametrdan idni olish
		id := c.Param("id")
		// Foydalanuvchini ro'yhatdan qidirish
		for _, user := range users {
			if strconv.Itoa(user.ID) == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		// Foydalanuvchi ma'lumotlari
		// user := User{
		// 	ID:   1,
		// 	Name: "Alice",
		// }

		// Malumotlarni Jsonda formatida qaytairsh
		// c.JSON(http.StatusOK, user)
	})

	r.GET("user", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080")
}
