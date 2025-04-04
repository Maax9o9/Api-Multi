package controllers

import (
	"Multi/src/house/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserHousesController struct {
	getUserHousesUseCase *application.GetHousesByUserIDUseCase
}

func NewGetUserHousesController(
	getUserHousesUseCase *application.GetHousesByUserIDUseCase,
) *GetUserHousesController {
	return &GetUserHousesController{
		getUserHousesUseCase: getUserHousesUseCase,
	}
}

func (c *GetUserHousesController) Handle(ctx *gin.Context) {
	userID := getUserIDFromContext(ctx)
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	houses, err := c.getUserHousesUseCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener casas del usuario", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, houses)
}
