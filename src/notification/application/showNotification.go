package application

import (
    "Multi/src/notification/domain"
    "Multi/src/notification/domain/entities"
)

type ShowAllNotificationsUseCase struct {
    repo domain.NotificationRepository
}

func NewShowAllNotificationsUseCase(repo domain.NotificationRepository) *ShowAllNotificationsUseCase {
    return &ShowAllNotificationsUseCase{
        repo: repo,
    }
}

func (uc *ShowAllNotificationsUseCase) GetAllNotifications() ([]entities.Notification, error) {
    return uc.repo.GetAll()
}