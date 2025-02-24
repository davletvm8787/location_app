package services

import (
	"Location-app/internal/models"
	"Location-app/internal/repository"
	"math"
	"sort"
)

type LocationService struct {
	repo *repository.LocationRepository
}

func NewLocationService(repo *repository.LocationRepository) *LocationService {
	return &LocationService{repo: repo}
}

func (ls *LocationService) CreateLocation(location models.Location) error {
	return ls.repo.Create(&location)
}

func (ls *LocationService) GetAllLocations() ([]models.Location, error) {
	return ls.repo.FindAll()
}

func (ls *LocationService) GetLocationByID(id uint) (models.Location, error) {
	return ls.repo.FindByID(id)
}

func (ls *LocationService) UpdateLocation(location models.Location) error {
	return ls.repo.Update(&location)
}

func (ls *LocationService) CalculateRoute(destLat, destLon float64) ([]models.Location, error) {
	locations, err := ls.repo.FindAll()
	if err != nil {
		return nil, err
	}
	sort.Slice(locations, func(i, j int) bool {
		distI := haversine(destLat, destLon, locations[i].Latitude, locations[i].Longitude)
		distJ := haversine(destLat, destLon, locations[j].Latitude, locations[j].Longitude)
		return distI < distJ
	})
	return locations, nil
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Радиус Земли в км
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
