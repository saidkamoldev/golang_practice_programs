package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	Age       int    `json:"age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var db *gorm.DB

func initDb() {
	//.env faylini yuklash
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// .nv dan DSN parametrlarni olish
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to mograte database:", err)
	}

}

func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func resetUserSequence() {
	db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1;")
}

func allDeleteUsers(c *gin.Context) {
	var user User

	if err := db.First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "database is clear"})
		return
	}

	db.Where("1=1").Delete(&User{})
	resetUserSequence()
	c.JSON(http.StatusOK, gin.H{"message": "All users deleted"})
}

func main() {

	initDb()
	r := gin.Default()

	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.DELETE("/users/:id", deleteUser)
	r.DELETE("/users/all/", allDeleteUsers)
	r.Run(":8080")

}
