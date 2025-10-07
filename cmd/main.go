package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/morodik/convoy/internal/db"
	"github.com/morodik/convoy/internal/handlers"
	"github.com/morodik/convoy/internal/repository"
	services "github.com/morodik/convoy/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить .env, используются системные переменные")
	}
	services.InitJWTKey()

	db.Init()
	userRepo := repository.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepo)
	profileHandler := handlers.NewProfileHandler(userService)

	convoyRepo := repository.NewConvoyRepository(db.DB)
	convoyService := services.NewConvoyService(convoyRepo)
	convoyHandler := handlers.NewConvoyHandler(convoyService)

	r := gin.Default()
	r.SetTrustedProxies(nil) // разрешить все локальные подключения (опционально, можно удалить)

	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Разрешаем все origins для тестирования (в продакшене укажите ["http://localhost:3000"])
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	r.POST("/register", services.Register)
	r.POST("/login", services.Login)
	r.POST("/logout", services.Logout)
	r.POST("/convoys", services.AuthMiddleware(), convoyHandler.CreateConvoy)

	r.GET("/profile", services.AuthMiddleware(), profileHandler.GetProfile)
	r.PUT("/profile", services.AuthMiddleware(), profileHandler.UpdateProfile)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
