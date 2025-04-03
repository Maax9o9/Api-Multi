package controllers

import (
    "Multi/src/house/application"
    "Multi/src/house/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AddHouseController struct {
    imageHandlerUseCase    *application.ImageHandlerUseCase
    locationHandlerUseCase *application.LocationHandlerUseCase
}

func NewAddHouseController(
    imageHandlerUseCase *application.ImageHandlerUseCase,
    locationHandlerUseCase *application.LocationHandlerUseCase,
) *AddHouseController {
    return &AddHouseController{
        imageHandlerUseCase:    imageHandlerUseCase,
        locationHandlerUseCase: locationHandlerUseCase,
    }
}

func (ac *AddHouseController) AddHouse(c *gin.Context) {
    var house entities.HouseProfile
    if err := c.ShouldBind(&house); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido", "details": err.Error()})
        return
    }

    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener la imagen", "details": err.Error()})
        return
    }

    src, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al abrir la imagen", "details": err.Error()})
        return
    }
    defer src.Close()

    imagePath, err := ac.imageHandlerUseCase.SaveImage(file.Filename, src)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la imagen", "details": err.Error()})
        return
    }
    house.Image = imagePath

    if err := ac.locationHandlerUseCase.SaveHouseLocation(&house); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la ubicación de la casa", "details": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Casa agregada exitosamente"})
}