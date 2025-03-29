package infrestructure

import (
    "Multi/src/user/application"
    "Multi/src/user/infrestructure/controllers"
)

func InitUser() (*controllers.CreateUserController, *controllers.ShowUserController, *controllers.LoginController) {
    userRepo := NewPostgres()

    createUserUseCase := application.NewCreateUserUseCase(userRepo)
    showUsersUseCase := application.NewShowUsersUseCase(userRepo)

    createUserController := controllers.NewCreateUserController(createUserUseCase)
    showUserController := controllers.NewShowUserController(showUsersUseCase)
    loginController := controllers.NewLoginController(userRepo)

    return createUserController, showUserController, loginController
}