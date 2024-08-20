package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Age       int
	CreateAt  time.Time
	UpdatedAt time.Time
}

func writeDb(db *gorm.DB) {

	user := User{Name: "Saidkamol", Email: "saidkamol.edu@gmail.com", Age: 20}
	result := db.Create(&user)

	if result.Error != nil {
		log.Fatal("Failed to create user:", result.Error)
	}

}

func readDb(db *gorm.DB) string {

	var user User
	db.First(&user, 1)
	db.First(&user, "email = ?", "saidkamol.edu@gmail.com")
	return user.Name
}

func main() {

	dsn := "host=localhost user=postgres password=123Uzb456 dbname=myfirst port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Fail to migrate database:", err)
	}

	writeDb(db)
	name := readDb(db)
	log.Println("User found:", name)

}
