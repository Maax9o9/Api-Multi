package controllers

import (
	"Multi/src/house/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetHouseByIDController struct {
	getHouseByIDUseCase *application.GetHouseByIDUseCase
}

func NewGetHouseByIDController(
	getHouseByIDUseCase *application.GetHouseByIDUseCase,
) *GetHouseByIDController {
	return &GetHouseByIDController{
		getHouseByIDUseCase: getHouseByIDUseCase,
	}
}

func (c *GetHouseByIDController) Handle(ctx *gin.Context) {
	houseIDStr := ctx.Param("id")
	houseID, err := strconv.Atoi(houseIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de casa inválido"})
		return
	}

	house, err := c.getHouseByIDUseCase.Execute(houseID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Casa no encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, house)
}
