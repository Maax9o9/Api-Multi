package controllers

import (
    "Multi/src/house/application"
    "Multi/src/house/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type LocationHandlerController struct {
    locationHandlerUseCase *application.LocationHandlerUseCase
}

func NewLocationHandlerController(locationHandlerUseCase *application.LocationHandlerUseCase) *LocationHandlerController {
    return &LocationHandlerController{
        locationHandlerUseCase: locationHandlerUseCase,
    }
}

func (lc *LocationHandlerController) SaveHouseLocation(c *gin.Context, house *entities.HouseProfile) error {
    if err := lc.locationHandlerUseCase.SaveHouseLocation(house); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la ubicación de la casa", "details": err.Error()})
        return err
    }

    return nil
}