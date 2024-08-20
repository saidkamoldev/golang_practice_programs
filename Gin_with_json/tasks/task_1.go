package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Userlar

type User struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

var users = []User{}

func main() {

	var newUser User
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		// var users User

		if c.Request.Method != http.MethodPost {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"error": "Method Not Allowed",
			})
			return
		}
		// if err := c.ShouldBindJSON(&newUser); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		newUser.ID = len(users) + 1
		users = append(users, newUser)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Users successful created",
		})

	})

	r.GET("/users", func(c *gin.Context) {

		// if c.Request.Method != http.MethodGet {
		// 	c.JSON(http.StatusMethodNotAllowed, gin.H{
		// 		"error": "Method Not Allowed",
		// 	})
		// 	return
		// }

		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080")

}
