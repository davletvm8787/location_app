package main

import (
	"Location-app/cmd/config"
	"Location-app/internal/handlers"
	"Location-app/internal/middleware"
	"Location-app/internal/models"
	"Location-app/internal/repository"
	"Location-app/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключаемся к базе данных
	db, err := gorm.Open(mysql.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	// Миграция модели
	if err := db.AutoMigrate(&models.Location{}); err != nil {
		panic("Ошибка миграции базы данных")
	}

	// Инициализируем слои
	repo := repository.NewLocationRepository(db)
	service := services.NewLocationService(repo)
	handler := handlers.NewLocationHandler(service)

	// Настраиваем Fiber
	app := fiber.New()

	// Применяем ограничение скорости
	app.Use(middleware.RateLimiter())

	// Маршруты
	api := app.Group("/api")
	api.Post("/locations", handler.AddLocation)
	api.Get("/locations", handler.ListLocations)
	api.Get("/locations/:id", handler.GetLocation)
	api.Put("/locations/:id", handler.EditLocation)
	api.Get("/route", handler.GetRoute)

	// Запускаем сервер
	if err := app.Listen(":3000"); err != nil {
		panic("ERROR START SERVER")
	}
}
