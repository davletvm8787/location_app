package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Модель Location
type Location struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"not null" json:"name" validate:"required"`
	Latitude    float64 `gorm:"not null" json:"latitude" validate:"required"`
	Longitude   float64 `gorm:"not null" json:"longitude" validate:"required"`
	MarkerColor string  `gorm:"not null;size:7" json:"marker_color" validate:"required,hexcolor"`
}

// Валидатор
var validate = validator.New()

// Настройка тестового приложения
func setupTestApp() (*fiber.App, *gorm.DB) {
	app := fiber.New()
	dsn := "app_user:secure_password@tcp(127.0.0.1:3306)/location_app_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Location{})
	return app, db
}

// Тест для Edit Location Endpoint
func TestEditLocation(t *testing.T) {
	app, db := setupTestApp()

	// Создаем тестовую локацию
	loc := Location{Name: "Old", Latitude: 55.7558, Longitude: 37.6173, MarkerColor: "#FF5733"}
	db.Create(&loc)

	app.Put("/locations/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var input Location
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
		}
		if err := validate.Struct(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		var loc Location
		if err := db.First(&loc, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "location not found"})
		}
		db.Model(&loc).Updates(input)
		return c.JSON(loc)
	})

	payload := `{"name": "Updated", "latitude": 55.7600, "longitude": 37.6200, "marker_color": "#33FF57"}`
	req := httptest.NewRequest("PUT", "/locations/1", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error testing request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	// Проверка валидации
	invalidPayload := `{"name": "", "latitude": 55.7600, "longitude": 37.6200, "marker_color": "#XYZ"}`
	req = httptest.NewRequest("PUT", "/locations/1", bytes.NewBuffer([]byte(invalidPayload)))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = app.Test(req)
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400 for invalid input, got %d", resp.StatusCode)
	}
}

// Тест для Location Details Endpoint
func TestLocationDetails(t *testing.T) {
	app, db := setupTestApp()

	// Создаем тестовую локацию
	loc := Location{Name: "Park", Latitude: 55.7581, Longitude: 37.6100, MarkerColor: "#5733FF"}
	db.Create(&loc)

	app.Get("/locations/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var loc Location
		if err := db.First(&loc, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "location not found"})
		}
		return c.JSON(loc)
	})

	req := httptest.NewRequest("GET", "/locations/1", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error testing request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	// Проверка ошибки для несуществующего ID
	req = httptest.NewRequest("GET", "/locations/999", nil)
	resp, _ = app.Test(req)
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404 for missing location, got %d", resp.StatusCode)
	}
}

// Тест для Routing Endpoint
func TestRouting(t *testing.T) {
	app, db := setupTestApp()

	// Создаем тестовые локации
	db.Create(&Location{Name: "A", Latitude: 55.7558, Longitude: 37.6173, MarkerColor: "#FF5733"})
	db.Create(&Location{Name: "B", Latitude: 55.7512, Longitude: 37.6231, MarkerColor: "#33FF57"})
	db.Create(&Location{Name: "C", Latitude: 55.7581, Longitude: 37.6100, MarkerColor: "#5733FF"})

	app.Get("/route", func(c *fiber.Ctx) error {
		destLat := c.QueryFloat("latitude", 0)
		destLon := c.QueryFloat("longitude", 0)
		if destLat == 0 || destLon == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "destination coordinates required"})
		}

		var locations []Location
		db.Find(&locations)

		// Простая сортировка по расстоянию (евклидово расстояние)
		type RoutePoint struct {
			Location
			Distance float64
		}
		var route []RoutePoint
		for _, loc := range locations {
			dist := ((loc.Latitude-destLat)*(loc.Latitude-destLat) + (loc.Longitude-destLon)*(loc.Longitude-destLon))
			route = append(route, RoutePoint{loc, dist})
		}
		// Сортировка вручную (для простоты)
		for i := 0; i < len(route)-1; i++ {
			for j := i + 1; j < len(route); j++ {
				if route[i].Distance > route[j].Distance {
					route[i], route[j] = route[j], route[i]
				}
			}
		}
		return c.JSON(route)
	})

	req := httptest.NewRequest("GET", "/route?latitude=55.7550&longitude=37.6150", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error testing request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	// Проверка валидации
	req = httptest.NewRequest("GET", "/route", nil)
	resp, _ = app.Test(req)
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400 for missing params, got %d", resp.StatusCode)
	}
}

func TestMain(m *testing.M) {
	// Здесь можно добавить запуск MySQL в контейнере для тестов, если нужно
	m.Run()
}
