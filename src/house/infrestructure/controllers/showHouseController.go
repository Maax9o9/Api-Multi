package controllers

import (
    "Multi/src/house/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ShowAllHousesController struct {
    showAllHousesUseCase *application.ShowAllHousesUseCase
}

func NewShowAllHousesController(showAllHousesUseCase *application.ShowAllHousesUseCase) *ShowAllHousesController {
    return &ShowAllHousesController{
        showAllHousesUseCase: showAllHousesUseCase,
    }
}

func (sc *ShowAllHousesController) GetAllHouses(c *gin.Context) {
    houses, err := sc.showAllHousesUseCase.GetAllHouses()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las casas", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, houses)
}