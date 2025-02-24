package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name      string  `json:"name" gorm:"not null" validate:"required"`
	Latitude  float64 `json:"latitude" gorm:"not null" validate:"required,latitude"`
	Longitude float64 `json:"longitude" gorm:"not null" validate:"required,longitude"`
	Color     string  `json:"color" gorm:"not null" validate:"required,hexcolor"`
}
