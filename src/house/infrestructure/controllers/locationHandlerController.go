package controllers

import (
	"Multi/src/house/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateLocationController struct {
	locationHandlerUseCase *application.LocationHandlerUseCase
	getHouseUseCase        *application.GetHouseByIDUseCase
	updateHouseUseCase     *application.UpdateHouseUseCase
}

func NewUpdateLocationController(
	locationHandlerUseCase *application.LocationHandlerUseCase,
	getHouseUseCase *application.GetHouseByIDUseCase,
	updateHouseUseCase *application.UpdateHouseUseCase,
) *UpdateLocationController {
	return &UpdateLocationController{
		locationHandlerUseCase: locationHandlerUseCase,
		getHouseUseCase:        getHouseUseCase,
		updateHouseUseCase:     updateHouseUseCase,
	}
}

func (c *UpdateLocationController) Handle(ctx *gin.Context) {
	// Obtener ID del usuario autenticado
	userID := getUserIDFromContext(ctx)
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	// Obtener ID de la casa a actualizar
	houseIDStr := ctx.Param("id")
	houseID, err := strconv.Atoi(houseIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de casa inválido"})
		return
	}

	// Obtener casa actual para verificar permisos
	house, err := c.getHouseUseCase.Execute(houseID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Casa no encontrada"})
		return
	}

	// Verificar que el usuario sea propietario de la casa
	if house.OwnerID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para actualizar esta casa"})
		return
	}

	// Obtener la nueva ubicación del cuerpo de la petición
	var request struct {
		UbicationGps string `json:"ubication_gps" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	// Actualizar la ubicación en la casa
	house.UbicationGps = request.UbicationGps

	// Procesar la ubicación con el caso de uso
	if err := c.locationHandlerUseCase.Execute(house); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error en la ubicación", "details": err.Error()})
		return
	}

	// Guardar los cambios
	if err := c.updateHouseUseCase.Execute(house); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la casa", "details": err.Error()})
		return
	}

	// Devolver la casa actualizada
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ubicación actualizada correctamente",
		"house":   house,
	})
}
