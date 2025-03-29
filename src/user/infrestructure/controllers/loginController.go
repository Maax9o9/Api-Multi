package controllers

import (
    "Multi/src/user/domain"
    "Multi/src/user/infrestructure/middlewares"
    "net/http"

    "github.com/gin-gonic/gin"
)

type LoginController struct {
    userRepo domain.UserRepository
}

func NewLoginController(userRepo domain.UserRepository) *LoginController {
    return &LoginController{
        userRepo: userRepo,
    }
}

func (lc *LoginController) Login(c *gin.Context) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido", "details": err.Error()})
        return
    }

    user, err := lc.userRepo.GetByUsername(credentials.Username)
    if err != nil || user.Password != credentials.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
        return
    }

    token, err := middlewares.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}