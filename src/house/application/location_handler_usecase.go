package application

import (
	"errors"
	"strings"

	"Multi/src/house/domain/entities"
)

type LocationHandlerUseCase struct{}

func NewLocationHandlerUseCase() *LocationHandlerUseCase {
	return &LocationHandlerUseCase{}
}

func (uc *LocationHandlerUseCase) Execute(house *entities.HouseProfile) error {
	if strings.TrimSpace(house.UbicationGps) == "" {
		return errors.New("la ubicación no puede estar vacía")
	}

	if len(house.UbicationGps) > 500 {
		house.UbicationGps = house.UbicationGps[:500]
	}

	house.UbicationGps = sanitizeLocation(house.UbicationGps)

	return nil
}

func sanitizeLocation(location string) string {

	location = strings.ReplaceAll(location, "<", "")
	location = strings.ReplaceAll(location, ">", "")
	location = strings.ReplaceAll(location, ";", "")
	location = strings.ReplaceAll(location, "'", "")
	location = strings.ReplaceAll(location, "\"", "")
	location = strings.ReplaceAll(location, "\\", "")

	return strings.TrimSpace(location)
}
