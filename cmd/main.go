package main

import (
	"log"
	"os"

	"github.com/morodik/convoy/internal/db"
	services "github.com/morodik/convoy/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить .env, используются системные переменные")
	}

	db.Init()

	r := gin.Default()

	r.POST("/register", services.Register)
	r.POST("/login", services.Login)
	r.POST("/logout", services.Logout)

	r.GET("/profile", services.AuthMiddleware(), func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		email, _ := c.Get("email")
		c.JSON(200, gin.H{
			"user_id": userID,
			"email":   email,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
