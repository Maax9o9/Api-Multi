package infrestructure

import (
    "Multi/src/user/application"
    "Multi/src/user/infrestructure/controllers"
)

func InitUser() (*controllers.CreateUserController, *controllers.ShowAllUsersController, *controllers.ShowUserByIDController, *controllers.LoginController) {
    userRepo := NewPostgres()

    createUserUseCase := application.NewCreateUserUseCase(userRepo)
    showAllUsersUseCase := application.NewShowAllUsersUseCase(userRepo)
    showUserByIDUseCase := application.NewShowUserByIDUseCase(userRepo)

    createUserController := controllers.NewCreateUserController(createUserUseCase)
    showAllUsersController := controllers.NewShowAllUsersController(showAllUsersUseCase)
    showUserByIDController := controllers.NewShowUserByIDController(showUserByIDUseCase)
    loginController := controllers.NewLoginController(userRepo)

    return createUserController, showAllUsersController, showUserByIDController, loginController
}