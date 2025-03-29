package controllers

import (
    "Multi/src/house/application"
    "Multi/src/house/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
)

type EditHouseController struct {
    editHouseUseCase *application.EditHouseUseCase
}

func NewEditHouseController(editHouseUseCase *application.EditHouseUseCase) *EditHouseController {
    return &EditHouseController{
        editHouseUseCase: editHouseUseCase,
    }
}

func (ec *EditHouseController) UpdateHouse(c *gin.Context) {
    var house entities.HouseProfile
    if err := c.ShouldBindJSON(&house); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido", "details": err.Error()})
        return
    }

    if err := ec.editHouseUseCase.UpdateHouse(&house); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la casa", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Casa actualizada exitosamente"})
}