package controllers

import (
    "Multi/src/user/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ShowAllUsersController struct {
    showAllUsersUseCase *application.ShowAllUsersUseCase
}

func NewShowAllUsersController(showAllUsersUseCase *application.ShowAllUsersUseCase) *ShowAllUsersController {
    return &ShowAllUsersController{
        showAllUsersUseCase: showAllUsersUseCase,
    }
}

func (sc *ShowAllUsersController) GetAllUsers(c *gin.Context) {
    users, err := sc.showAllUsersUseCase.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}