package controllers

import (
    "Multi/src/house/application"
    "Multi/src/house/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AddHouseController struct {
    addHouseUseCase *application.AddHouseUseCase
}

func NewAddHouseController(addHouseUseCase *application.AddHouseUseCase) *AddHouseController {
    return &AddHouseController{
        addHouseUseCase: addHouseUseCase,
    }
}

func (ac *AddHouseController) AddHouse(c *gin.Context) {
    var house entities.HouseProfile
    if err := c.ShouldBindJSON(&house); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido", "details": err.Error()})
        return
    }

    if err := ac.addHouseUseCase.AddHouse(&house); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar la casa", "details": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Casa agregada exitosamente"})
}