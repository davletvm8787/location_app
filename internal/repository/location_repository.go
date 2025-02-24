package repository

import (
	"Location-app/internal/models"

	"gorm.io/gorm"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

func (lr *LocationRepository) Create(location *models.Location) error {
	return lr.db.Create(location).Error
}

func (lr *LocationRepository) FindAll() ([]models.Location, error) {
	var locations []models.Location
	result := lr.db.Find(&locations)
	return locations, result.Error
}

func (lr *LocationRepository) FindByID(id uint) (models.Location, error) {
	var location models.Location
	result := lr.db.First(&location, id)
	return location, result.Error
}

func (lr *LocationRepository) Update(location *models.Location) error {
	return lr.db.Save(location).Error
}
