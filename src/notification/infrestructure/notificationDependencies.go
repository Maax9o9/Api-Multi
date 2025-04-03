package infrestructure

import (
    "Multi/src/notification/application"
    "Multi/src/notification/infrestructure/controllers"
)

func InitNotification() (*controllers.CreateNotificationController, *controllers.ShowAllNotificationsController, *controllers.ShowNotificationByIDController) {
    notificationRepo := NewPostgres()

    createNotificationUseCase := application.NewCreateNotificationUseCase(notificationRepo)
    showAllNotificationsUseCase := application.NewShowAllNotificationsUseCase(notificationRepo)
    showNotificationByIDUseCase := application.NewShowNotificationByIDUseCase(notificationRepo)

    createNotificationController := controllers.NewCreateNotificationController(createNotificationUseCase)
    showAllNotificationsController := controllers.NewShowAllNotificationsController(showAllNotificationsUseCase)
    showNotificationByIDController := controllers.NewShowNotificationByIDController(showNotificationByIDUseCase)

    return createNotificationController, showAllNotificationsController, showNotificationByIDController
}