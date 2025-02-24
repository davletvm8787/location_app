package validations

import (
	"Location-app/internal/models"
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type LocationValidator struct {
	validator *validator.Validate
}

func NewLocationValidator() *LocationValidator {
	v := validator.New()
	return &LocationValidator{validator: v}
}

func (lv *LocationValidator) Validate(location models.Location) error {
	if !isValidHexColor(location.Color) {
		return errors.New("Неверный формат hex-цвета")
	}
	return lv.validator.Struct(location)
}

func isValidHexColor(color string) bool {
	pattern := `^#[0-9A-Fa-f]{6}$`
	matched, _ := regexp.MatchString(pattern, color)
	return matched
}
