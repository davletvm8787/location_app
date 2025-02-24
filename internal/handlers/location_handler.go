package handlers

import (
	"Location-app/internal/models"
	"Location-app/internal/services"
	"Location-app/internal/validations"

	"github.com/gofiber/fiber/v2"
)

type LocationHandler struct {
	service *services.LocationService
}

func NewLocationHandler(service *services.LocationService) *LocationHandler {
	return &LocationHandler{service: service}
}

func (lh *LocationHandler) AddLocation(c *fiber.Ctx) error {
	var location models.Location
	if err := c.BodyParser(&location); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "HATALI  VERI"})
	}
	if err := validations.NewLocationValidator().Validate(location); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := lh.service.CreateLocation(location); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "CREATE ERROR"})
	}
	return c.Status(201).JSON(location)
}

func (lh *LocationHandler) ListLocations(c *fiber.Ctx) error {
	locations, err := lh.service.GetAllLocations()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "ERROR GET LIST"})
	}
	return c.JSON(locations)
}

func (lh *LocationHandler) GetLocation(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "HATALI ID"})
	}
	location, err := lh.service.GetLocationByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "LOKASYON BULUNUMADI"})
	}
	return c.JSON(location)
}

func (lh *LocationHandler) EditLocation(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "HATALI ID"})
	}
	var location models.Location
	if err := c.BodyParser(&location); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "HATALI  VERI"})
	}
	location.ID = uint(id)
	if err := validations.NewLocationValidator().Validate(location); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := lh.service.UpdateLocation(location); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Update error"})
	}
	return c.JSON(location)
}

func (lh *LocationHandler) GetRoute(c *fiber.Ctx) error {
	lat := c.QueryFloat("lat", 0) // Значение по умолчанию 0
	lon := c.QueryFloat("lon", 0) // Значение по умолчанию 0

	// Проверяем, что параметры переданы и не равны 0
	if lat == 0 || lon == 0 {
		// Дополнительно проверим, есть ли параметры в запросе
		if c.Query("lat") == "" || c.Query("lon") == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Широта и долгота обязательны"})
		}
		return c.Status(400).JSON(fiber.Map{"error": "Неверные значения широты или долготы"})
	}

	route, err := lh.service.CalculateRoute(lat, lon)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка построения маршрута"})
	}
	return c.JSON(route)
}
