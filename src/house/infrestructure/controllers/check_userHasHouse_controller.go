package controllers

import (
	"Multi/src/house/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckUserHasHouseController struct {
	checkUserHasHouseUseCase *application.CheckUserHasHouseUseCase
}

func NewCheckUserHasHouseController(
	checkUserHasHouseUseCase *application.CheckUserHasHouseUseCase,
) *CheckUserHasHouseController {
	return &CheckUserHasHouseController{
		checkUserHasHouseUseCase: checkUserHasHouseUseCase,
	}
}

func (c *CheckUserHasHouseController) Handle(ctx *gin.Context) {
	userID := getUserIDFromContext(ctx)
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	hasHouse, err := c.checkUserHasHouseUseCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar casas del usuario", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"has_house": hasHouse,
	})
}
