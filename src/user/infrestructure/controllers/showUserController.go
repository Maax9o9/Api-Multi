package controllers

import (
    "Multi/src/user/application"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ShowUserController struct {
    showUsersUseCase *application.ShowUsersUseCase
}

func NewShowUserController(showUsersUseCase *application.ShowUsersUseCase) *ShowUserController {
    return &ShowUserController{
        showUsersUseCase: showUsersUseCase,
    }
}

func (sc *ShowUserController) GetAllUsers(c *gin.Context) {
    users, err := sc.showUsersUseCase.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (sc *ShowUserController) GetUserByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido", "details": err.Error()})
        return
    }

    user, err := sc.showUsersUseCase.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}