package controllers

import (
    "Multi/src/house/application"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ShowHouseByIDController struct {
    showHouseByIDUseCase *application.ShowHouseByIDUseCase
}

func NewShowHouseByIDController(showHouseByIDUseCase *application.ShowHouseByIDUseCase) *ShowHouseByIDController {
    return &ShowHouseByIDController{
        showHouseByIDUseCase: showHouseByIDUseCase,
    }
}

func (sc *ShowHouseByIDController) GetHouseByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido", "details": err.Error()})
        return
    }

    house, err := sc.showHouseByIDUseCase.GetHouseByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Casa no encontrada", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, house)
}