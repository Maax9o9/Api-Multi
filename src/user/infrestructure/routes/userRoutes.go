package routes

import (
    "Multi/src/user/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, createUserController *controllers.CreateUserController, showUserController *controllers.ShowUserController, loginController *controllers.LoginController) {
    router.POST("/api/users", createUserController.CreateUser)
    router.GET("/api/users", showUserController.GetAllUsers)
    router.GET("/api/users/:id", showUserController.GetUserByID)
    router.POST("/api/users/login", loginController.Login)
}