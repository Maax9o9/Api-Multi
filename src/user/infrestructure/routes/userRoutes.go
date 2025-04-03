package routes

import (
    "Multi/src/user/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoutes(
    router *gin.Engine,
    createUserController *controllers.CreateUserController,
    showAllUsersController *controllers.ShowAllUsersController,
    showUserByIDController *controllers.ShowUserByIDController,
    loginController *controllers.LoginController,
) {
    router.POST("/api/users", createUserController.CreateUser)
    router.GET("/api/users", showAllUsersController.GetAllUsers)
    router.GET("/api/users/:id", showUserByIDController.GetUserByID)
    router.POST("/api/users/login", loginController.Login)
}