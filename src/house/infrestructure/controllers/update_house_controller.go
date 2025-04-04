package controllers

import (
	"Multi/src/house/application"
	"Multi/src/house/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateHouseController struct {
	updateHouseUseCase     *application.UpdateHouseUseCase
	getHouseByIDUseCase    *application.GetHouseByIDUseCase
	imageHandlerUseCase    *application.ImageHandlerUseCase
	locationHandlerUseCase *application.LocationHandlerUseCase
}

func NewUpdateHouseController(
	updateHouseUseCase *application.UpdateHouseUseCase,
	getHouseByIDUseCase *application.GetHouseByIDUseCase,
	imageHandlerUseCase *application.ImageHandlerUseCase,
	locationHandlerUseCase *application.LocationHandlerUseCase,
) *UpdateHouseController {
	return &UpdateHouseController{
		updateHouseUseCase:     updateHouseUseCase,
		getHouseByIDUseCase:    getHouseByIDUseCase,
		imageHandlerUseCase:    imageHandlerUseCase,
		locationHandlerUseCase: locationHandlerUseCase,
	}
}

func (c *UpdateHouseController) Handle(ctx *gin.Context) {
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

	// Obtener la casa actual para validar permisos y preservar campos que no se actualizan
	currentHouse, err := c.getHouseByIDUseCase.Execute(houseID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Casa no encontrada"})
		return
	}

	// Validar que el usuario sea propietario de la casa
	if currentHouse.OwnerID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para actualizar esta casa"})
		return
	}

	var house entities.HouseProfile
	if err := ctx.ShouldBind(&house); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido", "details": err.Error()})
		return
	}

	// Mantener el ID y otros campos fijos
	house.HouseID = houseID
	house.OwnerID = currentHouse.OwnerID
	house.CreatedAt = currentHouse.CreatedAt

	// Procesar imagen si está presente
	file, err := ctx.FormFile("image")
	if err == nil { // Solo procesar si hay una imagen
		src, err := file.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al abrir la imagen", "details": err.Error()})
			return
		}
		defer src.Close()

		imagePath, err := c.imageHandlerUseCase.Execute(file.Filename, src)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la imagen", "details": err.Error()})
			return
		}
		house.Image = imagePath
	} else {
		// Si no hay nueva imagen, mantener la anterior
		house.Image = currentHouse.Image
	}

	// Procesar ubicación
	if house.UbicationGps != currentHouse.UbicationGps {
		if err := c.locationHandlerUseCase.Execute(&house); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la ubicación", "details": err.Error()})
			return
		}
	}

	// Actualizar la casa
	if err := c.updateHouseUseCase.Execute(&house); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la casa", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Casa actualizada exitosamente",
		"house":   house,
	})
}
