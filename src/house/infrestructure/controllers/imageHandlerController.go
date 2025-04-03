package controllers

import (
    "Multi/src/house/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ImageHandlerController struct {
    imageHandlerUseCase *application.ImageHandlerUseCase
}

func NewImageHandlerController(imageHandlerUseCase *application.ImageHandlerUseCase) *ImageHandlerController {
    return &ImageHandlerController{
        imageHandlerUseCase: imageHandlerUseCase,
    }
}

func (ic *ImageHandlerController) SaveImage(c *gin.Context) (string, error) {
    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener la imagen", "details": err.Error()})
        return "", err
    }

    src, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al abrir la imagen", "details": err.Error()})
        return "", err
    }
    defer src.Close()

    imagePath, err := ic.imageHandlerUseCase.SaveImage(file.Filename, src)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la imagen", "details": err.Error()})
        return "", err
    }

    return imagePath, nil
}