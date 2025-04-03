package controllers

import (
    "Multi/src/house/application"
    "Multi/src/house/domain/entities"
    "net/http"
    "strconv"
    "io"

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
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido", "details": err.Error()})
        return
    }

    var house entities.HouseProfile
    if err := c.ShouldBindJSON(&house); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido", "details": err.Error()})
        return
    }

    house.HouseID = id

    file, err := c.FormFile("image")
    var fileReader io.Reader
    var fileName string

    if err == nil {
        src, err := file.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al abrir la imagen", "details": err.Error()})
            return
        }
        defer src.Close()

        fileReader = src
        fileName = file.Filename
    }

    if err := ec.editHouseUseCase.UpdateHouse(&house, fileName, fileReader); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la casa", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Casa actualizada exitosamente"})
}