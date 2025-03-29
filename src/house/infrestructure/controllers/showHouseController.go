package controllers

import (
    "Multi/src/house/application"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

type ShowHouseController struct {
    showHouseUseCase *application.ShowHouseUseCase
}

func NewShowHouseController(showHouseUseCase *application.ShowHouseUseCase) *ShowHouseController {
    return &ShowHouseController{
        showHouseUseCase: showHouseUseCase,
    }
}

func (sc *ShowHouseController) GetAllHouses(c *gin.Context) {
    houses, err := sc.showHouseUseCase.GetAllHouses()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las casas", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, houses)
}

func (sc *ShowHouseController) GetHouseByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido", "details": err.Error()})
        return
    }

    house, err := sc.showHouseUseCase.GetHouseByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Casa no encontrada", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, house)
}