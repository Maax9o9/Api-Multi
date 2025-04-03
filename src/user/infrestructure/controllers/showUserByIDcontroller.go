package controllers

import (
    "Multi/src/user/application"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ShowUserByIDController struct {
    showUserByIDUseCase *application.ShowUserByIDUseCase
}

func NewShowUserByIDController(showUserByIDUseCase *application.ShowUserByIDUseCase) *ShowUserByIDController {
    return &ShowUserByIDController{
        showUserByIDUseCase: showUserByIDUseCase,
    }
}

func (sc *ShowUserByIDController) GetUserByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido", "details": err.Error()})
        return
    }

    user, err := sc.showUserByIDUseCase.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}