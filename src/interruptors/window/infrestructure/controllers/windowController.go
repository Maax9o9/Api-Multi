package controllers

import (
    service "Multi/src/interruptors/window/application/services"
    "Multi/src/interruptors/window/domain/entities"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type WindowController struct {
    windowService *service.WindowService
}

func NewWindowController(windowService *service.WindowService) *WindowController {
    return &WindowController{
        windowService: windowService,
    }
}

func (wc *WindowController) GetAllWindowData(ctx *gin.Context) {
    data, err := wc.windowService.GetAllWindowData()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to retrieve window data",
            "details": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Window data retrieved successfully",
        "data":    data,
    })
}

func (wc *WindowController) GetWindowDataByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    // Cambia el nombre del método al correcto
    data, err := wc.windowService.GetWindowDataByID(id) // Asegúrate de que este método exista en WindowService
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error":   "Failed to retrieve window data",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Window data retrieved successfully",
        "data":    data,
    })
}

func (wc *WindowController) CreateWindowData(ctx *gin.Context) {
    var request entities.WindowSensor
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid request",
            "details": err.Error(),
        })
        return
    }

    err := wc.windowService.CreateWindowData(&request)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Failed to create window data",
            "details": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Window data created successfully",
    })
}