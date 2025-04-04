package controllers

import (
	"Multi/src/house/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteHouseController struct {
	deleteHouseUseCase  *application.DeleteHouseUseCase
	getHouseByIDUseCase *application.GetHouseByIDUseCase
}

func NewDeleteHouseController(
	deleteHouseUseCase *application.DeleteHouseUseCase,
	getHouseByIDUseCase *application.GetHouseByIDUseCase,
) *DeleteHouseController {
	return &DeleteHouseController{
		deleteHouseUseCase:  deleteHouseUseCase,
		getHouseByIDUseCase: getHouseByIDUseCase,
	}
}

func (c *DeleteHouseController) Handle(ctx *gin.Context) {
	userID := getUserIDFromContext(ctx)
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	houseIDStr := ctx.Param("id")
	houseID, err := strconv.Atoi(houseIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de casa inválido"})
		return
	}

	// Verificar que la casa exista y que el usuario sea propietario
	house, err := c.getHouseByIDUseCase.Execute(houseID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Casa no encontrada"})
		return
	}

	if house.OwnerID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para eliminar esta casa"})
		return
	}

	// Eliminar la casa
	if err := c.deleteHouseUseCase.Execute(houseID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la casa", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Casa eliminada exitosamente",
	})
}
