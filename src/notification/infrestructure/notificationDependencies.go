package infrestructure

import (
    "Multi/src/notification/application"
    "Multi/src/notification/infrestructure/controllers"
)

func InitNotification() (*controllers.CreateNotificationController, *controllers.ShowNotificationController) {
    notificationRepo := NewPostgres()

    createNotificationUseCase := application.NewCreateNotificationUseCase(notificationRepo)

    showNotificationUseCase := application.NewShowNotificationUseCase(notificationRepo)

    createNotificationController := controllers.NewCreateNotificationController(createNotificationUseCase)
    showNotificationController := controllers.NewShowNotificationController(showNotificationUseCase)

    return createNotificationController, showNotificationController
}