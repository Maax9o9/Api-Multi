package controllers

import (
	"Multi/src/house/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UploadImageController maneja las solicitudes para subir imágenes de casas
type UploadImageController struct {
	imageHandlerUseCase     *application.ImageHandlerUseCase
	updateHouseImageUseCase *application.UpdateHouseImageUseCase
}

// NewUploadImageController crea una nueva instancia de UploadImageController
func NewUploadImageController(
	imageHandlerUseCase *application.ImageHandlerUseCase,
	updateHouseUseCase *application.UpdateHouseImageUseCase,
) *UploadImageController {
	return &UploadImageController{
		imageHandlerUseCase:     imageHandlerUseCase,
		updateHouseImageUseCase: updateHouseUseCase,
	}
}

// Handle procesa la solicitud para subir una imagen de casa
func (c *UploadImageController) Handle(ctx *gin.Context) {
	// Verificar autenticación
	userID := getUserIDFromContext(ctx)
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	// Obtener ID de la casa
	houseIDStr := ctx.PostForm("house_id")
	if houseIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de casa no proporcionado"})
		return
	}

	houseID, err := strconv.Atoi(houseIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de casa inválido"})
		return
	}

	// Procesar la imagen
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No se proporcionó una imagen válida", "details": err.Error()})
		return
	}

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

	// Actualizar la casa con la ruta de la imagen - CORREGIDO
	err = c.updateHouseImageUseCase.Execute(houseID, imagePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la información de la casa", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Imagen subida con éxito",
		"image_path": imagePath,
	})
}
