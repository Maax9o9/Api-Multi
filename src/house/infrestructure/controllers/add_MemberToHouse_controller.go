package controllers

import (
	"Multi/src/house/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddMemberToHouseController struct {
	addMemberToHouseUseCase *application.AddMemberToHouseUseCase
}

func NewAddMemberToHouseController(
	addMemberToHouseUseCase *application.AddMemberToHouseUseCase,
) *AddMemberToHouseController {
	return &AddMemberToHouseController{
		addMemberToHouseUseCase: addMemberToHouseUseCase,
	}
}

func (c *AddMemberToHouseController) Handle(ctx *gin.Context) {
	userID := getUserIDFromContext(ctx)
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	var request struct {
		HouseID int    `json:"house_id" binding:"required"`
		UserID  int    `json:"user_id" binding:"required"`
		Role    string `json:"role" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	err := c.addMemberToHouseUseCase.Execute(request.HouseID, request.UserID, request.Role, userID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Error al añadir miembro", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Miembro añadido exitosamente",
	})
}
