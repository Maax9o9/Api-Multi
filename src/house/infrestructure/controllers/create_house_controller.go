package controllers

import (
	"Multi/src/house/application"
	"Multi/src/house/domain/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateHouseController struct {
	createHouseUseCase     *application.CreateHouseUseCase
	imageHandlerUseCase    *application.ImageHandlerUseCase
	locationHandlerUseCase *application.LocationHandlerUseCase
}

func NewCreateHouseController(
	createHouseUseCase *application.CreateHouseUseCase,
	imageHandlerUseCase *application.ImageHandlerUseCase,
	locationHandlerUseCase *application.LocationHandlerUseCase,
) *CreateHouseController {
	return &CreateHouseController{
		createHouseUseCase:     createHouseUseCase,
		imageHandlerUseCase:    imageHandlerUseCase,
		locationHandlerUseCase: locationHandlerUseCase,
	}
}

func (c *CreateHouseController) Handle(ctx *gin.Context) {

	userID := getUserIDFromContext(ctx)
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		log.Println("Error: Usuario no autenticado ", userID, ctx.Value("id"))
		return
	}

	var house entities.HouseProfile
	if err := ctx.ShouldBind(&house); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido", "details": err.Error()})
		return
	}

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
	}

	if err := c.locationHandlerUseCase.Execute(&house); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la ubicación", "details": err.Error()})
		return
	}

	houseID, err := c.createHouseUseCase.Execute(&house, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la casa", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{

		"message":  "Casa creada exitosamente",
		"house_id": houseID,
	})
}

func getUserIDFromContext(c *gin.Context) int {
	userIDValue, exists := c.Get("user_id")
	if !exists {
		return 0
	}

	userID, ok := userIDValue.(int)
	if !ok {
		if userIDFloat, ok := userIDValue.(float64); ok {
			return int(userIDFloat)
		}
		return 0
	}

	return userID
}
