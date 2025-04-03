package controllers

import (
    "Multi/src/user/application"
    "Multi/src/user/domain/entities"
    "net/http"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

type CreateUserController struct {
    createUserUseCase *application.CreateUserUseCase
}

func NewCreateUserController(createUserUseCase *application.CreateUserUseCase) *CreateUserController {
    return &CreateUserController{
        createUserUseCase: createUserUseCase,
    }
}

func (cc *CreateUserController) CreateUser(c *gin.Context) {
    var user entities.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido", "details": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear la contraseña", "details": err.Error()})
        return
    }
    user.Password = string(hashedPassword)

    if err := cc.createUserUseCase.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario", "details": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente"})
}